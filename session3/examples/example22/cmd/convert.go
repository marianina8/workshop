/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/marianina8/example22/utils"
	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts a JSON to YAML file",
	Run: func(cmd *cobra.Command, args []string) {
		// converts json to yaml by default
		file := cmd.Flag("input").Value.String()
		silence, _ := cmd.Flags().GetBool("silence")
		output := cmd.Flag("output").Value.String()

		if file == "" {
			fmt.Println("Please provide a JSON file to convert")
			return
		}
		yamlData, err := utils.J2Ytransform(file)
		if err != nil {
			fmt.Fprint(cmd.ErrOrStderr(), "Error converting JSON to YAML")
			os.Exit(1)
		}

		if !silence {
			fmt.Fprint(cmd.OutOrStdout(), string(yamlData))
		}

		if output != "" {
			err = ioutil.WriteFile(output, yamlData, 0644)
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Error writing YAML file: %v", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.PersistentFlags().StringP("file", "f", "", "JSON file to convert")
	convertCmd.PersistentFlags().BoolP("silence", "s", false, "Silence the output")
	convertCmd.PersistentFlags().StringP("output", "o", "", "Output file name")
}
