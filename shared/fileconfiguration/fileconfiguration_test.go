package fileconfiguration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ionos-cloud/sdk-go-bundle/shared"
)

func TestReadConfigFromFile(t *testing.T) {
	// Create a temporary config file
	tempFile, err := os.CreateTemp("", "config.yaml")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Write sample config data to the temp file
	configData := `
version: 1.0
currentProfile: testProfile
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
      - name: testProduct
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
	assert.Equal(t, 1.0, config.Version)
	assert.Equal(t, "testProfile", config.CurrentProfile)
	assert.Equal(t, "testUser", config.Profiles[0].Credentials.Username)
	assert.Equal(t, "testPass", config.Profiles[0].Credentials.Password)
	assert.Equal(t, "testToken", config.Profiles[0].Credentials.Token)
	assert.Equal(t, "testEnvironment", config.Environments[0].Name)
	assert.Equal(t, "testCertData", config.Environments[0].CertificateAuthData)
	assert.Equal(t, "testProduct", config.Environments[0].Products[0].Name)
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
      - name: testProduct
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
							{Location: "de/fra", Name: "https://dns.de-fra", SkipTLSVerify: false},
							{Location: "de/txl", Name: "https://dns.de-txl", SkipTLSVerify: false},
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
	// first endpoint in Product.Endpoints is the location-specific one, so fallback to GetProductOverrides:
	assert.Equal(t, "https://dns.de-fra", ep.Name)
}

func TestGetOverride_NotFound(t *testing.T) {
	cfg := makeTestConfig()
	// unknown product
	assert.Nil(t, cfg.GetOverride("unknown", ""))
	// known product but wrong location
	assert.Nil(t, cfg.GetOverride("dns", "wrong/location"))
}
