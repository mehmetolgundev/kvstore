/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kvstore",
	Short: "kvstore is a basic application that provides distributed key-value storage for trivial things.",
	Long:  `A distributed key value storage to handle basic operations or experience distributed applications nature.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
