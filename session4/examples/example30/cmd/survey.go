/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// surveyCmd represents the survey command
var surveyCmd = &cobra.Command{
	Use:   "survey",
	Short: "Prompts user for input",
	Run: func(cmd *cobra.Command, args []string) {
		var qs = []*survey.Question{
			{
				Name:      "name",
				Prompt:    &survey.Input{Message: "What is your name?"},
				Validate:  survey.Required,
				Transform: survey.Title,
			},
			{
				Name:     "age",
				Prompt:   &survey.Input{Message: "How old are you?"},
				Validate: survey.Required,
			},
			{
				Name: "programminglanguage",
				Prompt: &survey.Select{
					Message: "What is your favorite programming language?",
					Options: []string{"Go", "Java", "C++", "Python", "Ruby"},
					Default: "Go",
				},
			},
		}
		answers := struct {
			Name                string
			Age                 int
			ProgrammingLanguage string `survey:"programminglanguage"`
		}{}

		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("%s's favorite language is %s.", answers.Name, answers.ProgrammingLanguage)
	},
}

func init() {
	rootCmd.AddCommand(surveyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// surveyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// surveyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
