package app

import (
	"bytes"
	"io"
	"os"

	"github.com/gabefiori/gofp/internal/config"
	"github.com/gabefiori/gofp/internal/finder"
	"github.com/gabefiori/gofp/internal/ui"
	"github.com/mitchellh/go-homedir"
)

func Run(cfg *config.Config) error {
	home, err := homedir.Dir()

	if err != nil {
		return err
	}

	if !*cfg.ExpandResult {
		home = "~"
	}

	outputChan := make(chan string)

	go finder.Run(cfg.Sources, outputChan)

	result, err := ui.Run(outputChan)

	if err != nil {
		return err
	}

	if result != "" {
		_, err = io.Copy(os.Stdout, bytes.NewBufferString(home+result))
	}

	return err
}
