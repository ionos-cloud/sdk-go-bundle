package shared

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
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
    location: testLocation
    credentials:
      username: testUser
      password: testPass
      token: testToken
locations:
  - name: testLocation
    certificateAuthData: testCertData
    products:
      - name: testProduct
        endpoint: https://test.endpoint
        skipTlsVerify: true
`
	_, err = tempFile.Write([]byte(configData))
	assert.NoError(t, err)
	tempFile.Close()

	// Set the environment variable to point to the temp file
	os.Setenv(IonosFilePathEnvVar, tempFile.Name())
	defer os.Unsetenv(IonosFilePathEnvVar)

	// Call the function
	config, err := ReadConfigFromFile()
	assert.NoError(t, err)

	// Validate the loaded config
	assert.Equal(t, 1.0, config.Version)
	assert.Equal(t, "testProfile", config.CurrentProfile)
	assert.Equal(t, "testUser", config.Profiles[0].Credentials.Username)
	assert.Equal(t, "testPass", config.Profiles[0].Credentials.Password)
	assert.Equal(t, "testToken", config.Profiles[0].Credentials.Token)
	assert.Equal(t, "testLocation", config.Locations[0].Name)
	assert.Equal(t, "testCertData", config.Locations[0].CertificateAuthData)
	assert.Equal(t, "testProduct", config.Locations[0].Products[0].Name)
	assert.Equal(t, "https://test.endpoint", config.Locations[0].Products[0].Endpoint)
	assert.True(t, config.Locations[0].Products[0].SkipTLSVerify)
}

func TestDefaultLoadedConfigFileName(t *testing.T) {
	// Call the function
	fileName, err := DefaultLoadedConfigFileName()
	assert.NoError(t, err)
	assert.Contains(t, fileName, ".ionos")
	assert.Contains(t, fileName, "config")
}

func TestNewConfigurationFromLoaded_ValidConfig(t *testing.T) {
	lc := LoadedConfig{
		Version:        1.0,
		CurrentProfile: "testProfile",

		Profiles: []Profile{
			{
				Name:     "testProfile",
				Location: "testLocation",
				Credentials: struct {
					Username string `yaml:"username"`
					Password string `yaml:"password"`
					Token    string `yaml:"token"`
				}{
					Username: "testUser",
					Password: "testPass",
					Token:    "testToken",
				},
			},
		},
		Locations: []Location{
			{
				Name:                "testLocation",
				CertificateAuthData: "testCertData",
				Products: []struct {
					Name          string `yaml:"name"`
					Endpoint      string `yaml:"endpoint"`
					SkipTLSVerify bool   `yaml:"skipTlsVerify"`
				}{
					{
						Name:          "testProduct",
						Endpoint:      "https://test.endpoint",
						SkipTLSVerify: true,
					},
				},
			},
		},
	}

	config, err := NewConfigurationFromLoaded(lc, "testProduct")
	assert.NoError(t, err)
	assert.Equal(t, "testUser", config.Username)
	assert.Equal(t, "testPass", config.Password)
	assert.Equal(t, "testToken", config.Token)
	assert.Len(t, config.Servers, 1)
	assert.Equal(t, "https://test.endpoint", config.Servers[0].URL)
	assert.NotNil(t, config.HTTPClient)
	assert.True(t, config.HTTPClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
}

func TestNewConfigurationFromLoaded_NoProfile(t *testing.T) {
	lc := LoadedConfig{
		Version:        1.0,
		CurrentProfile: "",
		Profiles:       []Profile{},
		Locations: []Location{
			{
				Name:                "testLocation",
				CertificateAuthData: "testCertData",
				Products: []struct {
					Name          string `yaml:"name"`
					Endpoint      string `yaml:"endpoint"`
					SkipTLSVerify bool   `yaml:"skipTlsVerify"`
				}{
					{
						Name:          "testProduct",
						Endpoint:      "https://test.endpoint",
						SkipTLSVerify: true,
					},
				},
			},
		},
	}

	config, err := NewConfigurationFromLoaded(lc, "testProduct")
	assert.NoError(t, err)
	assert.Empty(t, config.Username)
	assert.Empty(t, config.Password)
	assert.Empty(t, config.Token)
	assert.Len(t, config.Servers, 1)
	assert.Equal(t, "https://test.endpoint", config.Servers[0].URL)
	assert.NotNil(t, config.HTTPClient)
	assert.True(t, config.HTTPClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
}

func TestNewConfigurationFromLoaded_NoMatchingProduct(t *testing.T) {
	lc := LoadedConfig{
		Version:        1.0,
		CurrentProfile: "testProfile",
		Profiles: []Profile{
			{
				Name:     "testProfile",
				Location: "testLocation",
				Credentials: struct {
					Username string `yaml:"username"`
					Password string `yaml:"password"`
					Token    string `yaml:"token"`
				}{
					Username: "testUser",
					Password: "testPass",
					Token:    "testToken",
				},
			},
		},
		Locations: []Location{
			{
				Name:                "testLocation",
				CertificateAuthData: "testCertData",
				Products: []struct {
					Name          string `yaml:"name"`
					Endpoint      string `yaml:"endpoint"`
					SkipTLSVerify bool   `yaml:"skipTlsVerify"`
				}{
					{
						Name:          "otherProduct",
						Endpoint:      "https://other.endpoint",
						SkipTLSVerify: false,
					},
				},
			},
		},
	}

	config, err := NewConfigurationFromLoaded(lc, "testProduct")
	assert.NoError(t, err)
	assert.Equal(t, "testUser", config.Username)
	assert.Equal(t, "testPass", config.Password)
	assert.Equal(t, "testToken", config.Token)
	assert.Empty(t, config.Servers)
	assert.Nil(t, config.HTTPClient)
}

func TestReadProfilesFromFile_ValidConfig(t *testing.T) {
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
locations:
  - name: testLocation
    certificateAuthData: testCertData
    products:
      - name: testProduct
        endpoint: https://test.endpoint
        skipTlsVerify: true
`
	_, err = tempFile.Write([]byte(configData))
	assert.NoError(t, err)
	tempFile.Close()

	os.Setenv(IonosFilePathEnvVar, tempFile.Name())
	defer os.Unsetenv(IonosFilePathEnvVar)

	profiles := ReadProfilesFromFile()
	assert.NotNil(t, profiles)
	assert.Len(t, profiles.Profiles, 1)
	assert.Equal(t, "testUser", profiles.Profiles[0].Credentials.Username)
	assert.Equal(t, "testPass", profiles.Profiles[0].Credentials.Password)
	assert.Equal(t, "testToken", profiles.Profiles[0].Credentials.Token)
}
