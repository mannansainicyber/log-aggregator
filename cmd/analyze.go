package cmd

import (
	"fmt"
	"log-ag/analyzer"
	"log-ag/db"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze logs for suspicious activity",
	Run: func(cmd *cobra.Command, args []string) {
		logs, err := db.Query("", "", "", 1000)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		alerts := analyzer.Analyze(logs)

		if len(alerts) == 0 {
			fmt.Println("No anomalies found.")
			return
		}

		for _, a := range alerts {
			fmt.Printf("[!] %-12s (%s)  %s\n", a.Type, a.Service, a.Message)
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
