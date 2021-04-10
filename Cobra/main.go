package main

import (
	"github.com/CristianQS/DoingStuffInGo/Cobra/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(commands.InitBeersCmd())
	rootCmd.Execute()
}
