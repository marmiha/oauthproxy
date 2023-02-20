package config

import (
	"fmt"
	"github.com/gume1a/oauth-proxy/internal/provider"
	"github.com/gume1a/oauth-proxy/pkg/identity"
	"gopkg.in/yaml.v3"
	"net/url"
	"os"
)

const (
	YamlConfigPath = ".template.oauthproxy.yaml"
	DefaultPort    = 8081
	DefaultHost    = "localhost"
)

type (
	Config struct {
		// The port to listen on.
		Port int `yaml:"port"`
		// The host to listen on.
		Host string `yaml:"host"`

		// The supported providers.
		Providers struct {
			// The supported providers.
			Supported []*SupportedProvider `yaml:"supported"`
			// The custom providers.
			Custom []*CustomProvider `yaml:"custom"`
		} `yaml:"providers"`
	}
	SupportedProvider struct {
		// The id of the provider.
		Id identity.ProviderId `yaml:"id"`
		// The client secret of the provider.
		ClientSecret string `yaml:"client_secret"`
	}
	CustomProvider struct {
		// The id of the provider.
		Id identity.ProviderId `yaml:"id"`
		// The token endpoint of the provider.
		TokenEndpoint string `yaml:"token_endpoint"`
		// The client secret of the provider.
		ClientSecret string `yaml:"client_secret"`
	}
)

var config *Config

func Get() Config {
	if config == nil {
		config = &Config{}

		data, err := os.ReadFile(YamlConfigPath)
		if err != nil && !os.IsNotExist(err) {
			panic(err)
		}

		// Set default values.
		if os.IsNotExist(err) {
			config.Host = DefaultHost
			config.Port = DefaultPort
			return *config
		}

		// Otherwise, parse the config file.
		err = yaml.Unmarshal(data, config)
		if err != nil {
			panic(err)
		}
	}
	return *config
}

func GetProviders() ([]provider.Provider, error) {
	cfg := Get()
	return cfg.GetProviders()
}

func (cfg *Config) GetProviders() ([]provider.Provider, error) {
	var providers []provider.Provider

	for _, entry := range cfg.Providers.Supported {
		// Get the client secret.
		secret, ok := os.LookupEnv(entry.ClientSecret)
		if !ok {
			return nil, fmt.Errorf("missing environment variable %v", entry.ClientSecret)
		}

		// Create the provider.
		var prov provider.Provider
		switch entry.Id {
		case identity.GITHUB:
			prov = provider.NewGithub(secret)
		case identity.GOOGLE:
			prov = provider.NewGoogle(secret)
		default:
			return nil, fmt.Errorf("unsupported provider %v", entry.Id)
		}
		providers = append(providers, prov)
	}

	for _, entry := range cfg.Providers.Custom {
		// Get the client secret.
		secret, ok := os.LookupEnv(entry.ClientSecret)
		if !ok {
			return nil, fmt.Errorf("missing environment variable %v", entry.ClientSecret)
		}

		// Create the provider.
		tokenURL, err := url.Parse(entry.TokenEndpoint)
		if err != nil {
			return nil, err
		}
		prov := provider.NewProvider(entry.Id, tokenURL, secret)
		providers = append(providers, prov)
	}

	return providers, nil
}
