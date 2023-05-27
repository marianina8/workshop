/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// yamlToJsonCmd represents the yamlToJson command
var yamlToJsonCmd = &cobra.Command{
	Use:   "yamlToJson",
	Short: "Converts yaml to json",
	Aliases: []string{"y2j"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yamlToJson called")
	},
}

func init() {
	convertCmd.AddCommand(yamlToJsonCmd)
}
