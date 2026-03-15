package cmd

import (
	"fmt"
	"log-ag/db"

	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show log counts by level and service",
	Run: func(cmd *cobra.Command, args []string) {
		stats, err := db.GetStats()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("─────────────────────────────────────────")
		fmt.Println("  BY LEVEL")
		fmt.Println("─────────────────────────────────────────")
		for _, s := range stats.ByLevel {
			fmt.Printf("  %-10s %d\n", s.Name, s.Count)
		}

		fmt.Println("─────────────────────────────────────────")
		fmt.Println("  BY SERVICE")
		fmt.Println("─────────────────────────────────────────")
		for _, s := range stats.ByService {
			fmt.Printf("  %-10s %d\n", s.Name, s.Count)
		}
		fmt.Println("─────────────────────────────────────────")
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
