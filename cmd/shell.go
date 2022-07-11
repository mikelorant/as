package cmd

import (
	"fmt"

	"github.com/mikelorant/asdfswitcher/internal/app"
	"github.com/spf13/cobra"
)

func NewShellCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shell",
		Short: "A brief description of your command",
		RunE: func(cmd *cobra.Command, args []string) error {
			switch len(args) {
			case 2:
				if err := app.Shell(
					app.WithPlugin(args[0]),
					app.WithVersion(args[1]),
				); err != nil {
					return fmt.Errorf("unable to launch shell for %v %v: %w", args[0], args[1], err)
				}
				return nil
			case 1:
				if err := app.Shell(
					app.WithPlugin(args[0]),
				); err != nil {
					return fmt.Errorf("unable to launch shell for %v: %w", args[0], err)
				}
				return nil
			}

			if err := app.Shell(); err != nil {
				return fmt.Errorf("unable to launch shell: %w", err)
			}
			return nil
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			switch len(args) {
			case 0:
				return app.GetPlugins(), cobra.ShellCompDirectiveNoFileComp
			case 1:
				return app.GetVersions(args[0]), cobra.ShellCompDirectiveNoFileComp
			}
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
	}

	return cmd
}
