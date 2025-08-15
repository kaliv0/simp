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
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Println(bashConfig)
		},
	}

	addShellConfigCmd = &cobra.Command{
		Use:   "add-config",
		Short: "Add integration config to .bashrc",
		Run: func(cmd *cobra.Command, _ []string) {
			f, err := os.OpenFile(path.Join(os.Getenv("HOME"), ".bashrc"), os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				//log.Println(err)
				//TODO
			}

			defer func(f *os.File) {
				err := f.Close()
				if err != nil {
					//log.Println(err)
					//TODO
				}
			}(f)

			_, err = f.WriteString(bashConfig)
			if err != nil {
				//TODO
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
	rootCmd.AddCommand(generateShellConfigCmd)
	rootCmd.AddCommand(addShellConfigCmd)

	rootCmd.AddCommand(completionCmd) // TODO: move to root?
}
