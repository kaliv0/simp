package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simp",
	Short: "Terminal-based clipboard manager",
}

func init() {
	rootCmd.Flags().BoolP(
		"help",
		"h",
		false,
		"Show help description for simp",
	)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
