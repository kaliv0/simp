package cmd

import (
	"os"
	"os/exec"

	"github.com/kaliv0/simp/pkg"
	"github.com/spf13/cobra"
)

var (
	startDaemonCmd = &cobra.Command{
		Use:   "start",
		Short: "Start clipboard manager", // (as daemon process)
		Run: func(cmd *cobra.Command, _ []string) {
			pkg.StopAllInstances()
			//if err != nil {
			//	//TODO
			//}

			err := exec.Command(os.Args[0], "run").Start()
			if err != nil {
				// TODO
			}
		},
	}

	runCmd = &cobra.Command{
		Use: "run",
		//Short:  "Run clipboard manager",
		Hidden: true,
		Run: func(cmd *cobra.Command, _ []string) {
			dbPath := pkg.GetDbPath()
			pkg.TrackClipboard(dbPath)
		},
	}

	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop clipboard manager",
		Run: func(cmd *cobra.Command, _ []string) {
			pkg.StopAllInstances()
			//if err != nil {
			//	//TODO
			//}
		},
	}
)

func init() {
	rootCmd.AddCommand(startDaemonCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(stopCmd)
}
