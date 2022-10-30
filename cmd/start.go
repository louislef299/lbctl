/*
Copyright © 2022 Louis Lefebvre <lefeb073@umn.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/louislef299/lbctl/internal/logserver"
	"github.com/spf13/cobra"
)

var port int

// logServerCmd represents the login command
var logServerCmd = &cobra.Command{
	Use:   "logserver",
	Short: "Start the distributed log server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		p := fmt.Sprintf(":%d", port)
		srv := logserver.NewHTTPServer(p)
		log.Fatal(srv.ListenAndServe())
	},
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start certain lbctl services",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.AddCommand(logServerCmd)
	logServerCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the log server on")
}
