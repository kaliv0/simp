package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	listHistoryCmd = &cobra.Command{
		Use:   "history",
		Short: "List clipboard history",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("history\n")
			// find db file handler (see runCmd),
			// read 'limit' && 'shouldPaste' flags
			// fetch history-to-be-displayed
			// load inside clipboard and (if shouldPaste) -> print history on command-line
		},
	}

	clearHistoryCmd = &cobra.Command{
		Use:   "clear",
		Short: "Clear clipboard history",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Print("clear\n")
			// find db file handler (see runCmd),
			// get repo handler -> reset db (delete table)
		},
	}
)

func init() {
	rootCmd.AddCommand(listHistoryCmd)
	listHistoryCmd.Flags().IntP(
		"limit",
		"l",
		20,
		"Limit the number of clipboard history items displayed",
	)
	listHistoryCmd.Flags().BoolP(
		"paste",
		"p",
		false,
		"Paste selected history item",
	)

	rootCmd.AddCommand(clearHistoryCmd)
}
