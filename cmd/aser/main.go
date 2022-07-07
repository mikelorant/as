package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/mikelorant/as/internal/app"
)

type InitCmd struct {
	Shell	string `arg:"positional"`
}

type ShellCmd struct {
	Plugin	string `arg:"positional"`
	Version	string `arg:"positional"`
}

type Args struct {
	Init	*InitCmd	`arg:"subcommand"`
	Shell	*ShellCmd	`arg:"subcommand"`
}

func main() {
	var args Args
	p := arg.MustParse(&args)

	switch {
	case args.Init != nil:
		app.Init(args.Init.Shell)
	case args.Shell != nil:
		if err := Shell(args.Shell); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
		}
	default:
		p.WriteHelp(os.Stdout)
	}
}

func Shell(a *ShellCmd) error {
	if a.Plugin == "" && a.Version == "" {
		if err := app.Shell(); err != nil {
			return fmt.Errorf("unable to launch shell: %w", err)
		}
		return nil
	}

	if a.Version == "" {
		if err := app.Shell(
			app.WithPlugin(a.Plugin),
		); err != nil {
			return fmt.Errorf("unable to launch shell for %v: %w", a.Plugin, err)
		}
		return nil
	}

	if err := app.Shell(
		app.WithPlugin(a.Plugin),
		app.WithVersion(a.Version),
	); err != nil {
		return fmt.Errorf("unable to launch shell for %v %v: %w", a.Plugin, a.Version, err)
	}

	return nil
}
