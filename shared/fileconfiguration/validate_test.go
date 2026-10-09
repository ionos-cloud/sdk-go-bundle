package fileconfiguration

import (
	"math"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ionos-cloud/sdk-go-bundle/shared/failover"
)

func TestNewRejectsSemanticallyInvalidConfig(t *testing.T) {
	tempFile, err := os.CreateTemp("", "config.yaml")
	assert.NoError(t, err)
	defer func() { _ = os.Remove(tempFile.Name()) }()

	// Syntactically valid YAML, but the cloud product mixes a global and a
	// location-based endpoint, which Validate rejects.
	configData := `
version: 1.0
currentProfile: testProfile
profiles:
  - name: testProfile
    environment: testEnvironment
    credentials:
      token: testToken
environments:
  - name: testEnvironment
    products:
      - name: cloud
        endpoints:
          - name: api.ionos.com
          - location: de/fra
            name: cloud.de-fra.ionos.com
`
	_, err = tempFile.Write([]byte(configData))
	assert.NoError(t, err)
	assert.NoError(t, tempFile.Close())

	config, err := New(tempFile.Name())
	assert.Nil(t, config)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "defines both global and location-based endpoints")
}

func float64Ptr(v float64) *float64 { return &v }

func TestValidate(t *testing.T) {
	validEndpoint := Endpoint{Name: "https://api.example.com"}
	validProduct := Product{Name: Cloud, Endpoints: []Endpoint{validEndpoint}}

	fullyValid := &FileConfig{
		CurrentProfile: "default",
		Profiles:       []Profile{{Name: "default", Environment: "prod"}},
		Environments:   []Environment{{Name: "prod", Products: []Product{validProduct}}},
		Failover:       &failover.Options{Strategy: failover.RoundRobin, MaxRetries: 3},
	}

	tests := []struct {
		name        string
		fc          *FileConfig
		wantErr     bool
		errContains []string
	}{
		{name: "nil config", fc: nil, wantErr: false},
		{name: "empty config", fc: &FileConfig{}, wantErr: false},
		{name: "fully valid config", fc: fullyValid, wantErr: false},

		// Profile and environment names and references.
		{
			name: "duplicate environment names",
			fc: &FileConfig{Environments: []Environment{
				{Name: "prod", Products: []Product{validProduct}},
				{Name: "prod", Products: []Product{validProduct}},
			}},
			wantErr:     true,
			errContains: []string{`duplicate environment name "prod"`},
		},
		{
			name: "duplicate profile names",
			fc: &FileConfig{Profiles: []Profile{
				{Name: "default"},
				{Name: "default"},
			}},
			wantErr:     true,
			errContains: []string{`duplicate profile name "default"`},
		},
		{
			name:        "currentProfile references unknown profile",
			fc:          &FileConfig{CurrentProfile: "unknown", Profiles: []Profile{{Name: "default"}}},
			wantErr:     true,
			errContains: []string{`currentProfile "unknown" does not match any defined profile`},
		},
		{
			name:        "profile references unknown environment",
			fc:          &FileConfig{Profiles: []Profile{{Name: "default", Environment: "unknown"}}},
			wantErr:     true,
			errContains: []string{`profile "default" references unknown environment "unknown"`},
		},
		{
			name:    "empty profile environment is allowed",
			fc:      &FileConfig{Profiles: []Profile{{Name: "default", Environment: ""}}},
			wantErr: false,
		},
		{
			name:    "empty currentProfile is allowed",
			fc:      &FileConfig{CurrentProfile: "", Profiles: []Profile{{Name: "default"}}},
			wantErr: false,
		},
		{
			name: "references are matched case-insensitively and trimmed",
			fc: &FileConfig{
				CurrentProfile: " DEFAULT ",
				Profiles:       []Profile{{Name: "default", Environment: "PROD"}},
				Environments:   []Environment{{Name: "prod", Products: []Product{validProduct}}},
			},
			wantErr: false,
		},

		// Products / endpoints.
		{
			name:        "unknown product",
			fc:          &FileConfig{Environments: []Environment{{Name: "prod", Products: []Product{{Name: "invalidproduct", Endpoints: []Endpoint{validEndpoint}}}}}},
			wantErr:     true,
			errContains: []string{`unknown product "invalidproduct" in environment "prod"`},
		},
		{
			name:        "product with no endpoints",
			fc:          &FileConfig{Environments: []Environment{{Name: "prod", Products: []Product{{Name: Cloud}}}}},
			wantErr:     true,
			errContains: []string{`product "cloud" in environment "prod" has no endpoints`},
		},
		{
			name:        "endpoint with empty name",
			fc:          &FileConfig{Environments: []Environment{{Name: "prod", Products: []Product{{Name: Cloud, Endpoints: []Endpoint{{Name: "  "}}}}}}},
			wantErr:     true,
			errContains: []string{`product "cloud" in environment "prod" has an endpoint with an empty name (position 1)`},
		},
		{
			name: "duplicate product in same environment",
			fc: &FileConfig{Environments: []Environment{{Name: "prod", Products: []Product{
				validProduct,
				validProduct,
			}}}},
			wantErr:     true,
			errContains: []string{`duplicate product "cloud" in environment "prod"`},
		},
		{
			name: "same product in different environments is allowed",
			fc: &FileConfig{Environments: []Environment{
				{Name: "prod", Products: []Product{validProduct}},
				{Name: "staging", Products: []Product{validProduct}},
			}},
			wantErr: false,
		},
		{
			name: "cloud product mixing global and location endpoints",
			fc: &FileConfig{Environments: []Environment{{Name: "prod", Products: []Product{{
				Name: Cloud,
				Endpoints: []Endpoint{
					{Name: "https://global.example.com"},
					{Name: "https://de-fra.example.com", Location: "de/fra"},
				},
			}}}}},
			wantErr:     true,
			errContains: []string{`product "cloud" in environment "prod" defines both global and location-based endpoints`},
		},
		{
			name: "cloud product with only global endpoints is allowed",
			fc: &FileConfig{Environments: []Environment{{Name: "prod", Products: []Product{{
				Name:      Cloud,
				Endpoints: []Endpoint{{Name: "https://g1.example.com"}, {Name: "https://g2.example.com"}},
			}}}}},
			wantErr: false,
		},
		{
			name: "cloud product with only location endpoints is allowed",
			fc: &FileConfig{Environments: []Environment{{Name: "prod", Products: []Product{{
				Name:      Cloud,
				Endpoints: []Endpoint{{Name: "https://de-fra.example.com", Location: "de/fra"}},
			}}}}},
			wantErr: false,
		},
		{
			name: "mixing across different products is allowed",
			fc: &FileConfig{Environments: []Environment{{Name: "prod", Products: []Product{
				{Name: Cloud, Endpoints: []Endpoint{{Name: "https://global.example.com"}}},
				{Name: Kafka, Endpoints: []Endpoint{{Name: "https://de-fra.example.com", Location: "de/fra"}}},
			}}}},
			wantErr: false,
		},

		// Failover.
		{
			name:        "invalid failover strategy",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: "random"}},
			wantErr:     true,
			errContains: []string{`invalid failover strategy "random", supported values are "none", "roundRobin" or an empty value`},
		},
		{
			name:        "invalid retryableMethods entry",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, RetryableMethods: []string{"G3T"}}},
			wantErr:     true,
			errContains: []string{`invalid failover retryableMethods entry "G3T"`},
		},
		{
			name:    "valid retryableMethods are accepted case-insensitively",
			fc:      &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, RetryableMethods: []string{"get", "POST", " delete "}}},
			wantErr: false,
		},
		{name: "failover strategy none", fc: &FileConfig{Failover: &failover.Options{Strategy: failover.None}}, wantErr: false},
		{name: "failover strategy roundRobin", fc: &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin}}, wantErr: false},
		{name: "failover strategy empty", fc: &FileConfig{Failover: &failover.Options{}}, wantErr: false},
		{
			name:        "negative failover maxRetries",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, MaxRetries: -1}},
			wantErr:     true,
			errContains: []string{"failover maxRetries must be >= 0, got -1"},
		},
		{
			name:        "negative backoff initialInterval",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{InitialInterval: -time.Second}}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.initialInterval must be >= 0, got -1s"},
		},
		{
			name:        "negative backoff maxInterval",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{MaxInterval: -time.Second}}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.maxInterval must be >= 0, got -1s"},
		},
		{
			name:        "zero backoff multiplier",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{Multiplier: float64Ptr(0)}}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.multiplier must be > 0, got 0"},
		},
		{
			name:        "negative backoff multiplier",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{Multiplier: float64Ptr(-2)}}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.multiplier must be > 0, got -2"},
		},
		{
			name:        "backoff randomizationFactor above 1",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{RandomizationFactor: float64Ptr(1.5)}}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.randomizationFactor must be between 0 and 1, got 1.5"},
		},
		{
			name:        "backoff randomizationFactor below 0",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{RandomizationFactor: float64Ptr(-0.1)}}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.randomizationFactor must be between 0 and 1, got -0.1"},
		},
		{
			name:        "backoff initialInterval above explicit maxInterval",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{InitialInterval: 500 * time.Millisecond, MaxInterval: 100 * time.Millisecond}}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.initialInterval (500ms) must not be greater than maxInterval (100ms)"},
		},
		{
			name:        "backoff initialInterval above default maxInterval",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{InitialInterval: 2 * time.Minute}}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.initialInterval (2m0s) must not be greater than maxInterval (1m0s)"},
		},
		{
			name:        "backoff explicit maxInterval below default initialInterval",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{MaxInterval: 100 * time.Millisecond}}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.initialInterval (500ms) must not be greater than maxInterval (100ms)"},
		},
		{
			name:    "backoff initialInterval equal to maxInterval",
			fc:      &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{InitialInterval: time.Second, MaxInterval: time.Second}}},
			wantErr: false,
		},
		{
			name:        "NaN backoff multiplier",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{Multiplier: float64Ptr(math.NaN())}}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.multiplier must be > 0, got NaN"},
		},
		{
			name:        "infinite backoff multiplier",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{Multiplier: float64Ptr(math.Inf(1))}}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.multiplier must be > 0, got +Inf"},
		},
		{
			name:        "NaN backoff randomizationFactor",
			fc:          &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{RandomizationFactor: float64Ptr(math.NaN())}}},
			wantErr:     true,
			errContains: []string{"failover exponentialBackoff.randomizationFactor must be between 0 and 1, got NaN"},
		},
		{
			name:    "valid backoff multiplier and randomizationFactor",
			fc:      &FileConfig{Failover: &failover.Options{Strategy: failover.RoundRobin, ExponentialBackoff: &failover.ExponentialBackoffOptions{Multiplier: float64Ptr(1.5), RandomizationFactor: float64Ptr(0.5)}}},
			wantErr: false,
		},

		// Aggregation: several problems reported together.
		{
			name: "multiple problems are aggregated",
			fc: &FileConfig{
				CurrentProfile: "unknown",
				Environments:   []Environment{{Name: "prod", Products: []Product{{Name: "invalidproduct"}}}},
			},
			wantErr: true,
			errContains: []string{
				`currentProfile "unknown" does not match any defined profile`,
				`unknown product "invalidproduct" in environment "prod"`,
				`product "invalidproduct" in environment "prod" has no endpoints`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fc.Validate()
			if !tt.wantErr {
				assert.NoError(t, err)
				return
			}
			assert.Error(t, err)
			for _, want := range tt.errContains {
				assert.Contains(t, err.Error(), want)
			}
		})
	}
}
