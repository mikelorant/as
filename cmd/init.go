package cmd

import (
	"fmt"
	_ "embed"

	"github.com/spf13/cobra"
)

var (
	//go:embed init_bash.sh
	BASHFunction string
	//go:embed init_zsh.zsh
	ZSHFunction string
)

func NewInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "A brief description of your command",
		ValidArgs: []string{"zsh", "bash"},
		Args: cobra.ExactValidArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "zsh":
 				fmt.Print(ZSHFunction)
			case "bash":
 				fmt.Print(BASHFunction)
			default:
 				fmt.Println("unsupported shell")
			}

			return nil
		},
	}

	return cmd
}
