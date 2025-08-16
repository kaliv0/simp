package cmd

import (
	"fmt"

	"github.com/kaliv0/simp/pkg"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var (
	listHistoryCmd = &cobra.Command{
		Use:   "history",
		Short: "List clipboard history",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("history\n")
			// find db file handler (see runCmd),
			dbPath := pkg.GetDbPath()

			// read 'limit' && 'shouldPaste' flags
			limit, err := cmd.Flags().GetInt("limit")
			if err != nil {
				// TODO
			}
			shouldPaste, err := cmd.Flags().GetBool("paste")
			if err != nil {
				// TODO
			}

			// fetch history-to-be-displayed -> extract function ?
			output, err := simp.ListHistory(dbPath, limit)
			if err != nil && err.Error() != "abort" {
				cmd.PrintErrln(err) // TODO: why cmd.Print?
			}

			// put output inside clipboard without pasting
			clipboard.Write(clipboard.FmtText, []byte(output))
			if shouldPaste {
				fmt.Print(string(clipboard.Read(clipboard.FmtText)))
			} else {
				clipboard.Read(clipboard.FmtText)
			}
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
	rootCmd.AddCommand(listHistoryCmd)
	rootCmd.AddCommand(clearHistoryCmd)
}
