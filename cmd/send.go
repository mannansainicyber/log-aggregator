package cmd

import (
	"fmt"
	"log-ag/db"

	"github.com/spf13/cobra"
)

var level string
var service string

var sendCmd = &cobra.Command{
	Use:   "send [message]",
	Short: "Send a log message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		message := args[0]

		err := db.Insert(level, service, message)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("[%s] (%s) %s\n", level, service, message)
	},
}

func init() {
	sendCmd.Flags().StringVarP(&level, "level", "l", "info", "Log level (info, warn, error)")
	sendCmd.Flags().StringVarP(&service, "service", "s", "default", "Service name")
	rootCmd.AddCommand(sendCmd)
}
