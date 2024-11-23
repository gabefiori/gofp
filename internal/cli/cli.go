package cli

import (
	"os"
	"runtime/debug"

	"github.com/gabefiori/gofp/internal/app"
	"github.com/gabefiori/gofp/internal/config"
	"github.com/urfave/cli/v2"
)

func Run() error {
	var cfgPath string
	var expandResult bool

	app := &cli.App{
		Name:        "Find Project",
		HelpName:    "gofp",
		Usage:       "Find projects",
		Description: "A simple tool for quickly finding projects.",
		Version:     getVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Load configuration from `file`",
				Value:       "~/.config/gofp/config.json",
				Destination: &cfgPath,
			},
			&cli.BoolFlag{
				Name:        "expand-output",
				Aliases:     []string{"eo"},
				Usage:       "Expand output",
				Value:       true,
				Destination: &expandResult,
			},
		},
		Action: func(ctx *cli.Context) error {
			cfg, err := config.Load(cfgPath)

			if ctx.IsSet("expand") {
				cfg.ExpandOutput = &expandResult
			} else if cfg.ExpandOutput == nil {
				cfg.ExpandOutput = &expandResult
			}

			if err != nil {
				return err
			}

			return app.Run(cfg)
		},
	}

	if err := app.Run(os.Args); err != nil {
		return err
	}

	return nil
}

func getVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}

	return "unknown"
}
