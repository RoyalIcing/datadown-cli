package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// cobra.OnInitialize(initConfig)
}

var rootCmd = &cobra.Command{
	Use: "datadown",
}

// Execute runs the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
