package shared

import (
	"crypto/tls"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testEndpoint = "https://test.endpoint"

func TestNewConfigurationFromOptions(t *testing.T) {
	tests := []struct {
		name           string
		clientOptions  ClientOptions
		expectedConfig *Configuration
	}{
		{
			name: "ValidOptions",
			clientOptions: ClientOptions{
				Endpoint:      testEndpoint,
				SkipTLSVerify: true,
				Certificate:   "",
				Credentials: Credentials{
					Username: "testUser",
					Password: "testPass",
					Token:    "testToken",
				},
			},
			expectedConfig: &Configuration{
				Username: "testUser",
				Password: "testPass",
				Token:    "testToken",
				Servers: ServerConfigurations{
					{
						URL:         testEndpoint,
						Description: "Production",
					},
				},
				HTTPClient: &http.Client{
					Transport: &http.Transport{
						TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
					},
				},
			},
		},
		{
			name: "EmptyEndpoint",
			clientOptions: ClientOptions{
				SkipTLSVerify: true,
				Certificate:   "",
				Credentials: Credentials{
					Username: "testUser",
					Password: "testPass",
					Token:    "testToken",
				},
			},
			expectedConfig: &Configuration{
				Username: "testUser",
				Password: "testPass",
				Token:    "testToken",
				Servers:  ServerConfigurations{},
				HTTPClient: &http.Client{
					Transport: &http.Transport{
						TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
					},
				},
			},
		},
		{
			name: "NoCredentials",
			clientOptions: ClientOptions{
				Endpoint:      testEndpoint,
				SkipTLSVerify: true,
				Certificate:   "",
			},
			expectedConfig: &Configuration{
				Username: "",
				Password: "",
				Token:    "",
				Servers: ServerConfigurations{
					{
						URL:         testEndpoint,
						Description: "Production",
					},
				},
				HTTPClient: &http.Client{
					Transport: &http.Transport{
						TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
					},
				},
			},
		},
		{
			name: "AddCertificate",
			clientOptions: ClientOptions{
				Endpoint:      testEndpoint,
				SkipTLSVerify: true,
				Certificate:   "testCertData",
				Credentials: Credentials{
					Username: "testUser",
					Password: "testPass",
					Token:    "testToken",
				},
			},
			expectedConfig: &Configuration{
				Username: "testUser",
				Password: "testPass",
				Token:    "testToken",
				Servers: ServerConfigurations{
					{
						URL:         testEndpoint,
						Description: "Production",
					},
				},
				HTTPClient: &http.Client{
					Transport: &http.Transport{
						TLSClientConfig: &tls.Config{
							InsecureSkipVerify: true,
							RootCAs:            AddCertsToClient("testCertData"),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfigurationFromOptions(tt.clientOptions)
			assert.Equal(t, tt.expectedConfig.Username, config.Username)
			assert.Equal(t, tt.expectedConfig.Password, config.Password)
			assert.Equal(t, tt.expectedConfig.Token, config.Token)
			assert.Equal(t, tt.expectedConfig.Servers, config.Servers)
			assert.NotNil(t, config.HTTPClient)
			assert.Equal(t, tt.expectedConfig.HTTPClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify,
				config.HTTPClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
			assert.True(t, config.HTTPClient.Transport.(*http.Transport).TLSClientConfig.RootCAs.Equal(tt.expectedConfig.HTTPClient.Transport.(*http.Transport).TLSClientConfig.RootCAs))
		})
	}
}

func TestNewConfigurationFromEnvWithObjectStorage(t *testing.T) {
	hostURL := "https://test.api.ionos.com"
	os.Setenv(IonosUsernameEnvVar, "testuser")
	os.Setenv(IonosPasswordEnvVar, "testpassword")
	os.Setenv(IonosTokenEnvVar, "testtoken")
	os.Setenv(IonosApiUrlEnvVar, hostURL)
	os.Setenv(IonosS3AccessKeyEnvVar, "testaccesskey")
	os.Setenv(IonosS3SecretKeyEnvVar, "testsecretkey")

	config := NewConfigurationFromEnv()

	if config.Username != "testuser" {
		t.Errorf("expected username to be 'testuser', got '%s'", config.Username)
	}
	if config.Password != "testpassword" {
		t.Errorf("expected password to be 'testpassword', got '%s'", config.Password)
	}
	if config.Token != "testtoken" {
		t.Errorf("expected token to be 'testtoken', got '%s'", config.Token)
	}
	if len(config.Servers) != 1 {
		t.Errorf("expected 1 server configuration, got %d", len(config.Servers))
	}
	if config.Servers[0].URL != hostURL {
		t.Errorf("expected host to be 'https://test.api.ionos.com', got '%s'", config.Host)
	}
	assert.NotNil(t, config.MiddlewareWithError, "MiddlewareWithError should not be nil as we have set objectstorage key and secret in env vars")
}

func TestConfigurationWithObjectStorage(t *testing.T) {
	username := "testuser"
	password := "testpassword"
	token := "testtoken"
	hostURL := "https://test.api.ionos.com"

	config := NewConfiguration(username, password, token, hostURL)

	s3AccessKey := "testaccesskey"
	s3SecretKey := "testsecretkey"
	options := ClientOptions{
		Credentials: Credentials{
			S3AccessKey: s3AccessKey,
			S3SecretKey: s3SecretKey,
		},
	}
	config = config.WithObjectStorage(options)

	assert.NotNil(t, config.MiddlewareWithError, "MiddlewareWithError should not be nil as we have set objectstorage key and secret in env vars")

	if config.Username != username {
		t.Errorf("Expected Username to be '%s', got '%s'", username, config.Username)
	}
	if config.Password != password {
		t.Errorf("Expected Password to be '%s', got '%s'", password, config.Password)
	}
	if config.Token != token {
		t.Errorf("Expected Token to be '%s', got '%s'", token, config.Token)
	}
	if config.Servers[0].URL != hostURL {
		t.Errorf("Expected Host URL to be '%s', got '%s'", hostURL, config.Servers[0].URL)
	}
}
