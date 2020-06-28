package cmd

import (
	"test-api/app/server"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serveCmd represents the run command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve serves the oms service",
	Long:  `Serve serves the oms service`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Println("run called")
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
