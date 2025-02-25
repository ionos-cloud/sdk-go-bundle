package fileconfiguration

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

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
