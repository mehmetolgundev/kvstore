/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Deletes the specified key.",
	Long: `DEL command is used to del a key from the kvstore. For example;
		kvstore del foo
		`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("del called")
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
