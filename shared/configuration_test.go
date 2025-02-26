package shared

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewConfigurationFromOptions(t *testing.T) {
	tests := []struct {
		name           string
		clientOptions  ClientOptions
		expectedConfig *Configuration
	}{
		{
			name: "ValidOptions",
			clientOptions: ClientOptions{
				Endpoint:      "https://test.endpoint",
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
						URL:         "https://test.endpoint",
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
				Endpoint:      "https://test.endpoint",
				SkipTLSVerify: true,
				Certificate:   "testCertData",
			},
			expectedConfig: &Configuration{
				Username: "",
				Password: "",
				Token:    "",
				Servers: ServerConfigurations{
					{
						URL:         "https://test.endpoint",
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
				Endpoint:      "https://test.endpoint",
				SkipTLSVerify: false,
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
						URL:         "https://test.endpoint",
						Description: "Production",
					},
				},
				HTTPClient: &http.Client{
					Transport: &http.Transport{
						TLSClientConfig: &tls.Config{},
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
			if tt.clientOptions.SkipTLSVerify {
				assert.True(t, config.HTTPClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
			}
		})
	}
}
