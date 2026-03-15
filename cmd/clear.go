package cmd

import (
	"fmt"
	"log-ag/db"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Delete all logs",
	Run: func(cmd *cobra.Command, args []string) {
		err := db.Clear()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Logs cleared!")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
