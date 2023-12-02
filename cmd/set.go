/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the given value to the given key",
	Long: `SET command is used to set a value to a key. For example;
		kvstore set foo bar
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("set called %v", args)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
