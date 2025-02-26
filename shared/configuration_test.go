package shared

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
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
