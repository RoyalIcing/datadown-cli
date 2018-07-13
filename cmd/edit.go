package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Run a datadown editor in the browser",
	Long:  `Run a datadown editor in the browser, by default on port 5775`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run a datadown editor in the browser")
	},
}
