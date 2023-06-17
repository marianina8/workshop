/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/marianina8/example27/utils"
	"github.com/spf13/cobra"
)

// yamlToJsonCmd represents the yamlToJson command
var yamlToJsonCmd = &cobra.Command{
	Use:     "yamlToJson",
	Short:   "Converts yaml to json",
	Aliases: []string{"y2j"},
	Run: func(cmd *cobra.Command, args []string) {
		file := cmd.Flag("file").Value.String()
		silence, _ := cmd.Flags().GetBool("silence")
		output := cmd.Flag("output").Value.String()

		if file == "" {
			fmt.Println("Please provide a YAML file to convert")
			return
		}
		jsonData, err := utils.Y2Jtransform(file)
		if err != nil {
			log.Fatalf("Error converting YAML to JSON: %v", err)
			os.Exit(1)
		}

		if !silence {
			fmt.Println(string(jsonData))
		}

		if output != "" {
			err = ioutil.WriteFile(output, jsonData, 0644)
			if err != nil {
				log.Fatalf("Error writing JSON file: %v", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	convertCmd.AddCommand(yamlToJsonCmd)
}
