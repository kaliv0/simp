package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

//go:embed scripts/.simprc
var bashConfig string

var (
	generateShellConfigCmd = &cobra.Command{
		Use:   "shell",
		Short: "Generate a shell integration script",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(bashConfig)
			shouldAdd, err := cmd.Flags().GetBool("add")
			if err != nil {
				//TODO
			}
			if shouldAdd {
				f, err := os.OpenFile(path.Join(os.Getenv("HOME"), ".bashrc"), os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					//TODO
				}

				defer func(f *os.File) {
					err := f.Close()
					if err != nil {
						//TODO
					}
				}(f)

				_, err = f.WriteString(fmt.Sprintf("\n%s", bashConfig))
				if err != nil {
					//TODO
				}
			}
		},
	}

	completionCmd = &cobra.Command{
		Use:    "completion",
		Hidden: true,
		// TODO: implement
	}
)

func init() {
	generateShellConfigCmd.Flags().BoolP(
		"add",
		"a",
		false,
		"Add integration config to .bashrc",
	)
	rootCmd.AddCommand(generateShellConfigCmd)
	rootCmd.AddCommand(completionCmd) // TODO: move to root?
}
