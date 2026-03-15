package cmd

import (
	"fmt"
	"log-ag/db"

	"github.com/spf13/cobra"
)

var searchLevel string
var searchService string
var since string
var limit int

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search through logs",
	Run: func(cmd *cobra.Command, args []string) {
		logs, err := db.Query(searchLevel, searchService, since, limit)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(logs) == 0 {
			fmt.Println("No logs found.")
			return
		}

		for _, l := range logs {
			fmt.Printf("%s  [%s]  (%s)  %s\n", l.Timestamp, l.Level, l.Service, l.Message)
		}
	},
}

func init() {
	searchCmd.Flags().StringVarP(&searchLevel, "level", "l", "", "Filter by level")
	searchCmd.Flags().StringVarP(&searchService, "service", "s", "", "Filter by service")
	searchCmd.Flags().StringVarP(&since, "since", "t", "", "Since duration e.g. 1 hour, 30 minutes")
	searchCmd.Flags().IntVarP(&limit, "limit", "n", 50, "Max number of logs to show")
	rootCmd.AddCommand(searchCmd)
}
