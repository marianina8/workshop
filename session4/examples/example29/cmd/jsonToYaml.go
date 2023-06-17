/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/marianina8/example29/utils"
	"github.com/spf13/cobra"
)

// jsonToYamlCmd represents the jsonToYaml command
var jsonToYamlCmd = &cobra.Command{
	Use:     "jsonToYaml",
	Short:   "Converts json to yaml",
	Aliases: []string{"j2y"},
	Run: func(cmd *cobra.Command, args []string) {
		file := cmd.Flag("file").Value.String()
		silence, _ := cmd.Flags().GetBool("silence")
		output := cmd.Flag("output").Value.String()

		if file == "" {
			fmt.Println("Please provide a JSON file to convert")
			return
		}
		yamlData, err := utils.J2Ytransform(file)
		if err != nil {
			log.Fatalf("Error converting JSON to YAML: %v", err)
			os.Exit(1)
		}

		if !silence {
			fmt.Println(string(yamlData))
		}

		if output != "" {
			err = ioutil.WriteFile(output, yamlData, 0644)
			if err != nil {
				log.Fatalf("Error writing YAML file: %v", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	convertCmd.AddCommand(jsonToYamlCmd)
}
