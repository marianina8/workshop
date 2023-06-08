/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Stores JSON data online",
	Long:  `Uploads JSON data to JSONBIN.io for later retrieval`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://api.jsonbin.io/v3/b"
		apiKey := "$2b$10$mwv0UXTT5hhwULEj8KehZuhT1Vgn8tg0FP6K9zr24TMWwqJDL95.."

		fileVal := cmd.Flag("file").Value.String()
		if fileVal == "" {
			fmt.Println("Please provide a JSON file to upload")
			return
		}
		file, err := os.Open(fileVal)
		if err != nil {
			fmt.Println("Erorr opening up file", err)
			return
		}
		defer file.Close()

		jsonBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading file", err)
			return
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			fmt.Println("Error creating request", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Master-Key", apiKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request", err)
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			fmt.Println("Error decoding response", err)
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error storing JSON data.  Received status: ", resp.Status)
			return
		} else {
			fmt.Println("File saved online at JSONBIN.io with id:", result["metadata"].(map[string]interface{})["id"])
			return
		}

	},
}

func init() {
	storeCmd.Flags().StringP("file", "f", "", "JSON file to upload")
	rootCmd.AddCommand(storeCmd)
}
