package app

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed init_bash.sh
	BASHFunction string
	//go:embed init_zsh.zsh
	ZSHFunction string
)

func Init(shell string) {
	switch shell {
	case "zsh":
		fmt.Print(ZSHFunction)
	case "bash":
		fmt.Print(BASHFunction)
	default:
		fmt.Println("unsupported shell")
	}
}
