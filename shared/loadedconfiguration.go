package shared

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type ClientOverrideOptions struct {
	Endpoint      string
	SkipTLSVerify bool
	Certificate   string
	Credentials   Credentials
}

type Credentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Token    string `yaml:"token"`
}

type Product struct {
	// Name is the name of the product
	Name string `yaml:"name"`
	// Endpoint is the endpoint that will be overridden
	Endpoint string `yaml:"endpoint"`
	// SkipTLSVerify skips tls verification. Not recommended for production
	SkipTLSVerify bool `yaml:"skipTlsVerify"`
}

// Location is a struct that represents a location
type Location struct {
	Name string `yaml:"name"`
	// CertificateAuthData
	CertificateAuthData string `yaml:"certificateAuthData,omitempty"`
	// Products is a list of ionos products for which we will override endpoint, tls verification
	Products []Product `yaml:"products"`
}

// Profiles wrapper to read only the profiles from the config file
type Profiles struct {
	// CurrentProfile active profile for configuration
	CurrentProfile string `yaml:"currentProfile"`
	// Profiles
	Profiles []Profile `yaml:"profiles"`
}

// Profile is a struct that represents a profile and it's Credentials
type Profile struct {
	Name        string `yaml:"name"`
	Location    string `yaml:"location"`
	Credentials Credentials
}

// LoadedConfig is a struct that represents the loaded configuration
type LoadedConfig struct {
	// Version of the configuration
	Version float64 `yaml:"version"`
	// CurrentProfile active profile for configuration
	CurrentProfile string `yaml:"currentProfile"`
	// Profiles list of profiles
	Profiles []Profile `yaml:"profiles"`
	// Locations list of locations
	Locations []Location `yaml:"locations"`
}

// DefaultLoadedConfigFileName returns the default file path for the loaded configuration
func DefaultLoadedConfigFileName() (string, error) {
	home := ""
	var err error
	if home, err = os.UserHomeDir(); err != nil {
		return home, err
	}
	if home == "" {
		return home, fmt.Errorf("could not determine home directory")
	}
	return filepath.Join(home, ".ionos", "config"), nil

}

func findConfigFile() string {
	filePath := ""
	filePath = os.Getenv(IonosFilePathEnvVar)
	var err error
	if filePath == "" {
		if filePath, err = DefaultLoadedConfigFileName(); err != nil {
			return ""
		}
	}
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		return ""
	}
	return filePath
}

// ReadProfilesFromFile reads profiles from yaml file, loads it into a struct and returns it
func ReadProfilesFromFile() *Profiles {
	filePath := findConfigFile()
	if filePath == "" {
		return nil
	}
	var content []byte
	var err error
	if content, err = os.ReadFile(filePath); err != nil {
		return nil
	}
	loadedProfiles := &Profiles{}
	_ = yaml.Unmarshal(content, loadedProfiles)
	return loadedProfiles
}

// ReadConfigFromFile reads yaml file, loads it into a struct and returns it
func ReadConfigFromFile() (*LoadedConfig, error) {
	filePath := ""
	filePath = os.Getenv(IonosFilePathEnvVar)
	var err error
	if filePath == "" {
		if filePath, err = DefaultLoadedConfigFileName(); err != nil {
			return nil, err
		}
	}
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s does not exist", filePath)
	}
	var content []byte
	if content, err = os.ReadFile(filePath); err != nil {
		return nil, err
	}
	if len(content) == 0 {
		return nil, fmt.Errorf("file %s is empty", filePath)
	}
	loadedConfig := &LoadedConfig{}
	if err = yaml.Unmarshal(content, loadedConfig); err != nil {
		return nil, fmt.Errorf("while unmarshalling file %s, %w", filePath, err)
	}
	return loadedConfig, nil
}

func (loadedConfig *LoadedConfig) GetCurrentProfile() *Profile {
	currentProfile := loadedConfig.CurrentProfile
	if currentProfile == "" {
		currentProfile = os.Getenv(IonosCurrentProfileEnvVar)
	}
	if currentProfile == "" {
		log.Printf("[WARN] no current profile set")
	}
	for _, profile := range loadedConfig.Profiles {
		if profile.Name == loadedConfig.CurrentProfile {
			return &profile
		}
	}
	return nil
}

// NewConfigurationFromLoaded creates a new configuration from a loaded
// configuration for a specific product
func NewConfigurationFromLoaded(lc LoadedConfig, sdkProductName string) (*Configuration, error) {
	config := &Configuration{}
	getCredentialsFromEnv := false
	// don't read Credentials from file, if they are set in the environment
	if os.Getenv(IonosUsernameEnvVar) != "" || os.Getenv(IonosTokenEnvVar) != "" {
		getCredentialsFromEnv = true
		config = NewConfigurationFromEnv()
	}
	profileLocation := ""
	if !getCredentialsFromEnv {
		if currentProfile := lc.GetCurrentProfile(); currentProfile != nil {
			config.Username = currentProfile.Credentials.Username
			config.Password = currentProfile.Credentials.Password
			config.Token = currentProfile.Credentials.Token
			profileLocation = currentProfile.Location
		}
	}
	lc.processLocations(config, profileLocation, sdkProductName)
	return config, nil
}

// processLocations - overrides http client and server configuration based on user location
func (loadedConfig *LoadedConfig) processLocations(config *Configuration, profileLocation, sdkProductName string) {
	for _, location := range loadedConfig.Locations {
		for _, product := range location.Products {
			if product.Name != sdkProductName {
				continue
			}
			switch profileLocation {
			// todo check if this is correct, should it actually read all locations if profileLocation is empty?
			case "":
				fallthrough
			case location.Name:
				config.Servers = append(config.Servers, ServerConfiguration{
					URL:         product.Endpoint,
					Description: "endpoint from loaded file",
				})
				config.HTTPClient = &http.Client{
					Transport: &http.Transport{
						TLSClientConfig: &tls.Config{
							InsecureSkipVerify: product.SkipTLSVerify,
						},
					},
				}
				AddCertsToClient(config.HTTPClient, location.CertificateAuthData)
			default:
				continue
			}
		}
	}
}

func (loadedConfig *LoadedConfig) GetLocationForCurrentProfile() string {
	if currentProfile := loadedConfig.GetCurrentProfile(); currentProfile != nil {
		return currentProfile.Location
	}
	return ""
}

func (loadedConfig *LoadedConfig) GetOverridesFor(productName string) *Product {
	if loadedConfig == nil {
		return nil
	}
	for _, location := range loadedConfig.Locations {
		for _, product := range location.Products {
			if product.Name == productName {
				return &product
			}
		}
	}
	return nil
}

// AddCertsToClient adds certificates to the http client
func AddCertsToClient(httpClient *http.Client, authorityData string) {
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	if ok := rootCAs.AppendCertsFromPEM([]byte(authorityData)); !ok {
		log.Println("No certs appended, using system certs only")
	}
	httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: rootCAs,
			},
		},
	}
}
