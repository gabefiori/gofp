package config

import (
	"os"

	"github.com/gabefiori/gofp/internal/finder"
	"github.com/goccy/go-json"
	"github.com/mitchellh/go-homedir"
)

type Config struct {
	Sources      []finder.Source `json:"sources"`
	ExpandOutput *bool           `json:"expand_output"`
}

func Load(path string) (*Config, error) {
	path, err := homedir.Expand(path)

	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	var cfg Config

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
