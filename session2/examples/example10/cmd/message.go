/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// messageCmd represents the message command
var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "Prints a hello world message",
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()
		if name == "" {
			fmt.Println("Hello World")
		} else {
			fmt.Printf("Hello World, %s\n", name)
		}
	},
}

func init() {
	messageCmd.Flags().StringP("name", "n", "", "Name to use in the message")
	rootCmd.AddCommand(messageCmd)
}
