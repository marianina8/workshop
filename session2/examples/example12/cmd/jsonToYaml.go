/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// jsonToYamlCmd represents the jsonToYaml command
var jsonToYamlCmd = &cobra.Command{
	Use:     "jsonToYaml",
	Short:   "Converts json to yaml",
	Aliases: []string{"j2y"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jsonToYaml called")
	},
}

func init() {
	convertCmd.AddCommand(jsonToYamlCmd)
}
