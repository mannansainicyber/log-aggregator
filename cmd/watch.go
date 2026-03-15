package cmd

import (
	"fmt"
	"log-ag/db"
	"time"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch logs in real time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Watching logs... (Ctrl+C to stop)")
		fmt.Println("─────────────────────────────────────────")

		lastID := 0

		// get the current last ID so we only show NEW logs
		logs, err := db.Query("", "", "", 1)
		if err == nil && len(logs) > 0 {
			lastID = logs[0].ID
		}

		for {
			newLogs, err := db.QueryAfter(lastID)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			for _, l := range newLogs {
				fmt.Printf("%s  [%s]  (%s)  %s\n", l.Timestamp, l.Level, l.Service, l.Message)
				lastID = l.ID
			}

			time.Sleep(1 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
