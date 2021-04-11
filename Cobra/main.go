package main

import (
	"Cobra/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "commands"}
	rootCmd.AddCommand(commands.InitBeersCmd())
	rootCmd.AddCommand(commands.InitStoreCmd())
	rootCmd.Execute()
}
