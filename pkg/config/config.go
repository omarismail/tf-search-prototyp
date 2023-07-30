package config

import (
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Root struct {
	NamedProviders []NamedProvider `hcl:"provider,block"`
	Searches       []Search        `hcl:"search,block"`
	Providers      map[string]*Provider
}

type NamedProvider struct {
	Name     string   `hcl:"name,label"`
	Provider Provider `hcl:",remain"`
}

type Provider struct {
	AccessKey string `hcl:"access_key,optional"`
	SecretKey string `hcl:"secret_key,optional"`
}

type Search struct {
	Name     string       `hcl:"name,label"`
	Provider string       `hcl:"provider"`
	Queries  []TypedQuery `hcl:"query,block"`
}

type TypedQuery struct {
	Type   string `hcl:"type,label"`
	Region string `hcl:"region,optional"`
	Query  Query  `hcl:",remain"`
}

type Query struct {
	Category string   `hcl:"category"`
	Tags     []string `hcl:"tags"`
}

type AwsTag struct {
	Key   string
	Value string
}

func LoadConfig(filename string) (*Root, error) {
	var config Root
	if !strings.HasSuffix(filename, ".hcl") {
		filename += ".hcl"
	}
	fullpath := filename

	err := hclsimple.DecodeFile(fullpath, nil, &config)
	if err != nil {
		return nil, err
	}

	// Create a map from provider name to Provider
	config.Providers = make(map[string]*Provider)
	for _, np := range config.NamedProviders {
		p := np.Provider
		if p.AccessKey == "" {
			p.AccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
		}

		if p.SecretKey == "" {
			p.SecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
		}

		config.Providers[np.Name] = &p
	}

	return &config, nil
}
