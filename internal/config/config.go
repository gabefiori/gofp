package config

import (
	"os"

	"github.com/gabefiori/gofp/internal/finder"
	"github.com/goccy/go-json"
	"github.com/mitchellh/go-homedir"
)

// Config represents the configuration structure for the application.
type Config struct {
	// List of sources to be used by the finder
	Sources []finder.Source `json:"sources"`

	// Optional flag to indicate if output should be expanded
	// Useful to hide the user's home directory
	ExpandOutput *bool `json:"expand_output"`

	// Flag to indicate if measurement should be performed
	Measure bool
}

// Load reads the configuration from a JSON file at the specified path.
func Load(path string) (*Config, error) {
	path, err := homedir.Expand(path)

	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var cfg Config

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
