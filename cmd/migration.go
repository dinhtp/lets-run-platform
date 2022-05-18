package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Run platform service migration command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run platform service serve command called")
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}
