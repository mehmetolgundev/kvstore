/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a key",
	Long: `GET command is used to get a key from the kvstore. For example;
		kvstore get foo
		`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("get called %v", args)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

}
