package cli

import (
	"os"
	"runtime/debug"

	"github.com/gabefiori/gofp/internal/app"
	"github.com/gabefiori/gofp/internal/config"
	"github.com/urfave/cli/v2"
)

// Run initializes and executes the command-line interface (CLI) application.
func Run() error {
	var (
		cfgPath      string
		expandResult bool
		measure      bool
	)

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
			&cli.BoolFlag{
				Name:        "measure",
				Aliases:     []string{"m"},
				Usage:       "Measure performance (time taken and number of items)",
				Value:       false,
				Destination: &measure,
			},
		},

		Action: func(ctx *cli.Context) error {
			cfg, err := config.Load(cfgPath)

			if err != nil {
				return err
			}

			cfg.Measure = measure

			if ctx.IsSet("expand") {
				cfg.ExpandOutput = &expandResult
			} else if cfg.ExpandOutput == nil {
				cfg.ExpandOutput = &expandResult
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
