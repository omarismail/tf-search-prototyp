package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tfsearch",
	Short: "tfsearch is a CLI tool for searching terraform",
	Long:  `A longer description should go here.`,
}

func Execute() error {
	return rootCmd.Execute()
}
