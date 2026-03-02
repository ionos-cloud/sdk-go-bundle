package compute_test

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	compute "github.com/ionos-cloud/sdk-go-bundle/products/compute/v2"
	"github.com/ionos-cloud/sdk-go-bundle/shared"
)

// minimalBackoff keeps tests fast without real waits.
var minimalBackoff = &shared.ExponentialBackoffOptions{
	InitialInterval: 1 * time.Millisecond,
	MaxInterval:     10 * time.Millisecond,
}

// datacenterResponse returns a minimal valid JSON body for a POST /datacenters response.
func datacenterResponse(id, name string) []byte {
	b, _ := json.Marshal(map[string]any{
		"id":   id,
		"type": "datacenter",
		"href": "https://api.ionos.com/cloudapi/v6/datacenters/" + id,
		"properties": map[string]any{
			"name":     name,
			"location": "de/fra",
		},
	})
	return b
}

// deadServerURL returns a URL whose port is immediately closed after allocation,
// so any connection attempt produces a "connection refused" *net.OpError.
func deadServerURL(t *testing.T) string {
	t.Helper()
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("deadServerURL: could not allocate port: %v", err)
	}
	addr := l.Addr().String()
	l.Close()
	return "http://" + addr
}

// failoverConfig builds a Configuration with two servers and failover enabled for POST.
func failoverConfig(server1URL, server2URL string, extraOpts func(*shared.FailoverOptions)) *shared.Configuration {
	fo := &shared.FailoverOptions{
		Strategy:           shared.FailoverRoundRobin,
		RetryableMethods:   []string{http.MethodPost},
		ExponentialBackoff: minimalBackoff,
	}
	if extraOpts != nil {
		extraOpts(fo)
	}
	return &shared.Configuration{
		Username: "test-user",
		Password: "test-pass",
		Servers: shared.ServerConfigurations{
			{URL: server1URL},
			{URL: server2URL},
		},
		Failover:   fo,
		MaxRetries: 2,
		// Use an explicit Transport so NewAPIClient does not fall back to
		// DeepCopy, which loses HTTPClient and falls back to http.DefaultClient,
		// stacking FailoverRoundTripper instances across tests.
		HTTPClient: &http.Client{Transport: http.DefaultTransport},
	}
}

// TestCreateDatacenter_FailoverOnNetworkError verifies that when the first
// configured server refuses connections, the SDK fails over to the second
// server and the datacenter is created successfully.
func TestCreateDatacenter_FailoverOnNetworkError(t *testing.T) {
	server2Calls := 0
	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server2Calls++
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(datacenterResponse("dc-failover-net", "failover-net-test"))
	}))
	defer server2.Close()

	cfg := failoverConfig(deadServerURL(t), server2.URL, nil)
	client := compute.NewAPIClient(cfg)

	props := compute.NewDatacenterPropertiesPost("de/fra")
	props.SetName("failover-net-test")

	dc, _, err := client.DataCentersApi.
		DatacentersPost(context.Background()).
		Datacenter(*compute.NewDatacenterPost(*props)).
		Execute()

	if err != nil {
		t.Fatalf("expected success after network-error failover, got: %v", err)
	}
	if server2Calls != 1 {
		t.Fatalf("expected server2 to be called exactly once, got %d", server2Calls)
	}
	if dc.GetId() != "dc-failover-net" {
		t.Fatalf("unexpected datacenter id: %q", dc.GetId())
	}
}

// TestCreateDatacenter_FailoverOnServiceUnavailable verifies that when the
// first server returns 503, the SDK fails over to the second server and the
// datacenter is created successfully.
func TestCreateDatacenter_FailoverOnServiceUnavailable(t *testing.T) {
	server1Calls := 0
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server1Calls++
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(`{"messages":[{"errorCode":"503","message":"service unavailable"}]}`))
	}))
	defer server1.Close()

	server2Calls := 0
	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server2Calls++
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(datacenterResponse("dc-failover-503", "failover-503-test"))
	}))
	defer server2.Close()

	cfg := failoverConfig(server1.URL, server2.URL, func(fo *shared.FailoverOptions) {
		fo.FailoverOnStatusCodes = []int{http.StatusServiceUnavailable}
	})
	client := compute.NewAPIClient(cfg)

	props := compute.NewDatacenterPropertiesPost("de/fra")
	props.SetName("failover-503-test")

	dc, _, err := client.DataCentersApi.
		DatacentersPost(context.Background()).
		Datacenter(*compute.NewDatacenterPost(*props)).
		Execute()

	if err != nil {
		t.Fatalf("expected success after 503 failover, got: %v", err)
	}
	if server1Calls != 1 {
		t.Fatalf("expected server1 to be called exactly once, got %d", server1Calls)
	}
	if server2Calls != 1 {
		t.Fatalf("expected server2 to be called exactly once, got %d", server2Calls)
	}
	if dc.GetId() != "dc-failover-503" {
		t.Fatalf("unexpected datacenter id: %q", dc.GetId())
	}
}
