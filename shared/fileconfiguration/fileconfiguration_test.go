package fileconfiguration

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ionos-cloud/sdk-go-bundle/shared"
	"github.com/ionos-cloud/sdk-go-bundle/shared/failover"
)

func TestReadConfigFromFile(t *testing.T) {
	// Create a temporary config file
	tempFile, err := os.CreateTemp("", "config.yaml")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Write sample config data to the temp file
	configData := `
version: 1.0
currentProfile: TESTPROFILE
profiles:
  - name: testProfile
    environment: testEnvironment
    credentials:
      username: testUser
      password: testPass
      token: testToken
environments:
  - name: testEnvironment
    certificateAuthData: testCertData
    products:
      - name: mariadb
        endpoints: 
          - location: de/fra
            name: mariadb.de-fra.ionos.com
            skipTlsVerify: false
          - location: de/txl
            name: mariadb.de-txl.ionos.com
            certificateAuthData: "certauthdata"
            skipTlsVerify: true
`
	_, err = tempFile.Write([]byte(configData))
	assert.NoError(t, err)
	tempFile.Close()

	// Set the environment variable to point to the temp file
	os.Setenv(shared.IonosFilePathEnvVar, tempFile.Name())
	defer os.Unsetenv(shared.IonosFilePathEnvVar)

	// Call the function
	config, err := NewFromEnv()
	assert.NoError(t, err)

	// Validate the loaded config
	assert.Equal(t, Version(1.0), config.Version)
	assert.Equal(t, strings.ToLower("testProfile"), strings.ToLower(config.CurrentProfile))
	assert.Equal(t, "testUser", config.Profiles[0].Credentials.Username)
	assert.Equal(t, "testPass", config.Profiles[0].Credentials.Password)
	assert.Equal(t, "testToken", config.Profiles[0].Credentials.Token)
	assert.Equal(t, "testEnvironment", config.Environments[0].Name)
	assert.Equal(t, "testCertData", config.Environments[0].CertificateAuthData)
	assert.Equal(t, "mariadb", config.Environments[0].Products[0].Name)
}

func TestDefaultLoadedConfigFileName(t *testing.T) {
	// Call the function
	fileName, err := DefaultConfigFileName()
	assert.NoError(t, err)
	assert.Contains(t, fileName, ".ionos")
	assert.Contains(t, fileName, "config")
}

func TestReadProfilesFromConfigFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "config.yaml")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	configData := `
version: 1.0
currentProfile: testProfile
profiles:
  - name: testProfile
    location: testLocation
    credentials:
      username: testUser
      password: testPass
      token: testToken
environments:
  - name: testEnvironment
    certificateAuthData: testCertData
    products:
      - name: mariadb
        endpoints: 
          - location: de/fra
            name: mariadb.de-fra.ionos.com
            skipTlsVerify: false
          - location: de/txl
            name: mariadb.de-txl.ionos.com
            certificateAuthData: "certauthdata"
            skipTlsVerify: true
`
	_, err = tempFile.Write([]byte(configData))
	assert.NoError(t, err)
	tempFile.Close()

	os.Setenv(shared.IonosFilePathEnvVar, tempFile.Name())
	defer os.Unsetenv(shared.IonosFilePathEnvVar)

	profiles := ReadProfilesFromFile()
	assert.NotNil(t, profiles)
	assert.Len(t, profiles.Profiles, 1)
	assert.Equal(t, "testUser", profiles.Profiles[0].Credentials.Username)
	assert.Equal(t, "testPass", profiles.Profiles[0].Credentials.Password)
	assert.Equal(t, "testToken", profiles.Profiles[0].Credentials.Token)
	assert.Equal(t, "testProfile", profiles.CurrentProfile)
}

func makeTestConfig() *FileConfig {
	return &FileConfig{
		CurrentProfile: "alice",
		Profiles: []Profile{
			{Name: "alice", Environment: "prod"},
			{Name: "bob", Environment: "dev"},
		},
		Environments: []Environment{
			{
				Name: "prod",
				Products: []Product{
					{
						Name: "psql",
						Endpoints: []Endpoint{
							{Name: "https://global.psql", SkipTLSVerify: false},
						},
					},
					{
						Name: "dns",
						Endpoints: []Endpoint{
							{Location: "", Name: "https://global.dns", SkipTLSVerify: false},
							{Location: "de/fra", Name: "https://dns.de-fra", SkipTLSVerify: false},
							{Location: "de/txl", Name: "https://dns.de-txl", SkipTLSVerify: false},
						},
					},
					{
						Name: Cloud,
						Endpoints: []Endpoint{
							{Location: "de/fra", Name: "https://cloud.de-fra", SkipTLSVerify: false},
							{Location: "de/txl", Name: "https://cloud.de-txl", SkipTLSVerify: false},
							{Location: "", Name: "https://cloud.global-1", SkipTLSVerify: false},
							{Location: "", Name: "https://cloud.global-2", SkipTLSVerify: true},
						},
					},
				},
			},
			{
				Name: "dev",
				Products: []Product{
					{
						Name: "psql",
						Endpoints: []Endpoint{
							{Name: "https://dev.psql", SkipTLSVerify: true},
						},
					},
				},
			},
		},
	}
}

func TestGetProfileNames(t *testing.T) {
	cfg := makeTestConfig()
	names := cfg.GetProfileNames()
	assert.ElementsMatch(t, []string{"alice", "bob"}, names, "should return both profile names")
}

func TestGetEnvironmentNames(t *testing.T) {
	cfg := makeTestConfig()
	envs := cfg.GetEnvironmentNames()
	assert.ElementsMatch(t, []string{"prod", "dev"}, envs, "should return both environment names")
}

func TestGetOverride_LocationMatch(t *testing.T) {
	cfg := makeTestConfig()

	var ep *Endpoint
	ep = cfg.GetOverride("dns", "de/fra")
	assert.NotNil(t, ep)
	assert.Equal(t, "https://dns.de-fra", ep.Name)
	assert.False(t, ep.SkipTLSVerify)

	ep = cfg.GetOverride("dns", "de/txl")
	assert.NotNil(t, ep)
	assert.Equal(t, "https://dns.de-txl", ep.Name)
	assert.False(t, ep.SkipTLSVerify)
}

func TestGetOverride_FallbackToGlobal(t *testing.T) {
	cfg := makeTestConfig()
	ep := cfg.GetOverride("psql", "")
	assert.NotNil(t, ep, "fallback global endpoint should be returned")
	assert.Equal(t, "https://global.psql", ep.Name)
	assert.False(t, ep.SkipTLSVerify)
}

func TestGetOverride_GlobalWhenLocationEmpty(t *testing.T) {
	cfg := makeTestConfig()
	ep := cfg.GetOverride("dns", "")
	assert.NotNil(t, ep)
	assert.Equal(t, "https://global.dns", ep.Name)
}

func TestGetOverride_NotFound(t *testing.T) {
	cfg := makeTestConfig()
	// unknown product
	assert.Nil(t, cfg.GetOverride("unknown", ""))
	// known product but wrong location, fallback to global endpoint (first in endpoint list), so should not be nil
	assert.NotNil(t, cfg.GetOverride("dns", "wrong/location"))
}

func TestFilterOverrides(t *testing.T) {
	cfg := makeTestConfig()
	ep := cfg.FilterOverrides(
		Cloud, func(endpoint Endpoint) bool {
			return endpoint.SkipTLSVerify == true
		},
	)
	assert.NotNil(t, ep)
	assert.Equal(t, 1, len(ep))
	assert.Equal(t, "https://cloud.global-2", ep[0].Name)
	assert.Equal(t, true, ep[0].SkipTLSVerify)
}

func TestFilterGlobalOverrides(t *testing.T) {
	cfg := makeTestConfig()
	ep := cfg.FilterGlobalOverrides(Cloud)
	assert.NotNil(t, ep)
	assert.Equal(t, 2, len(ep))
	assert.Equal(t, "", ep[0].Location)
	assert.Equal(t, "", ep[1].Location)
}

func TestFilterLocationOverrides(t *testing.T) {
	cfg := makeTestConfig()
	ep := cfg.FilterLocationOverrides(Cloud)
	assert.NotNil(t, ep)
	assert.Equal(t, 2, len(ep))
	assert.Equal(t, "de/fra", ep[0].Location)
	assert.Equal(t, "de/txl", ep[1].Location)
}

func TestGetLocationOverridesWithGlobalFallback_LocationMatch(t *testing.T) {
	cfg := makeTestConfig()
	ep := cfg.GetLocationOverridesWithGlobalFallback(Cloud, "de/fra")
	assert.NotNil(t, ep)
	assert.Equal(t, "https://cloud.de-fra", ep.Name)
	assert.Equal(t, "de/fra", ep.Location)
}

func TestGetLocationOverridesWithGlobalFallback_GlobalMatch(t *testing.T) {
	cfg := makeTestConfig()
	ep := cfg.GetLocationOverridesWithGlobalFallback(Cloud, "")
	assert.NotNil(t, ep)
	assert.Equal(t, "https://cloud.global-1", ep.Name)
	assert.Equal(t, "", ep.Location)
}

func TestGetLocationOverridesWithGlobalFallback_LocationNotFound_GlobalFallback(t *testing.T) {
	cfg := makeTestConfig()
	ep := cfg.GetLocationOverridesWithGlobalFallback(Cloud, "us/las")
	assert.NotNil(t, ep)
	assert.Equal(t, "https://cloud.global-1", ep.Name)
	assert.Equal(t, "", ep.Location)
}

func TestFailoverOptionsDeserializedFromYAML(t *testing.T) {
	tempFile, err := os.CreateTemp("", "config-failover-*.yaml")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

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
    products: []
failover:
  strategy: roundRobin
  retryableMethods:
    - GET
    - PUT
  retryOnTimeout: true
  failoverOnStatusCodes:
    - 502
    - 503
`
	_, err = tempFile.Write([]byte(configData))
	assert.NoError(t, err)
	tempFile.Close()

	cfg, err := New(tempFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, cfg.Failover)
	assert.Equal(t, failover.RoundRobin, cfg.Failover.Strategy)
	assert.Equal(t, []string{"GET", "PUT"}, cfg.Failover.RetryableMethods)
	assert.True(t, cfg.Failover.RetryOnTimeout)
	assert.Equal(t, []int{502, 503}, cfg.Failover.FailoverOnStatusCodes)
}

func TestFailoverOptionsNilWhenNotInFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "config-no-failover-*.yaml")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

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
    products: []
`
	_, err = tempFile.Write([]byte(configData))
	assert.NoError(t, err)
	tempFile.Close()

	cfg, err := New(tempFile.Name())
	assert.NoError(t, err)
	assert.Nil(t, cfg.Failover)
	assert.Nil(t, cfg.GetFailoverOptions())
}

func TestGetFailoverOptions(t *testing.T) {
	fo := &failover.Options{Strategy: failover.RoundRobin}
	fileCfg := &FileConfig{Failover: fo}
	assert.Equal(t, fo, fileCfg.GetFailoverOptions())

	var nilCfg *FileConfig
	assert.Nil(t, nilCfg.GetFailoverOptions())
}

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
