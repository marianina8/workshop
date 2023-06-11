/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/mum4k/termdash/widgets/textinput"
	"github.com/spf13/cobra"
)

// dashboardCmd represents the dashboard command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "An example dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		terminalLayer, err := tcell.New(tcell.ColorMode(terminalapi.ColorMode256), tcell.ClearStyle(cell.ColorYellow, cell.ColorNavy))
		if err != nil {
			fmt.Println(err)
		}
		defer terminalLayer.Close()
		value := make(chan string)
		rollingText, err := text.New(text.RollContent())
		if err != nil {
			panic(err)
		}
		go rollText(rollingText, value)
		leftContainer := container.Left(
			container.Border(linestyle.Light),
			container.PlaceWidget(rollingText),
		)
		input, err := textinput.New(
			textinput.Label("Input: "),
			textinput.MaxWidthCells(20),
			textinput.PlaceHolder("Enter text here"),
		)
		if err != nil {
			panic(err)
		}
		submitButton, err := button.New("Submit", func() error {
			inputValue := input.ReadAndClear()
			if inputValue != "" {
				value <- inputValue
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
		clearButton, err := button.New("Clear", func() error {
			rollingText.Reset()
			return nil
		})
		if err != nil {
			panic(err)
		}
		rightContainer := container.Right(
			container.SplitHorizontal(
				container.Top(
					container.Border(linestyle.Light),
					container.PlaceWidget(input),
				),
				container.Bottom(
					container.SplitVertical(
						container.Left(
							container.Border(linestyle.Light),
							container.PlaceWidget(submitButton),
						),
						container.Right(
							container.Border(linestyle.Light),
							container.PlaceWidget(clearButton),
						),
					),
				),
			),
		)
		containerLayer, _ := container.New(
			terminalLayer,
			container.SplitVertical(
				leftContainer,
				rightContainer,
				container.SplitPercent(60),
			),
		)

		ctx, cancel := context.WithCancel(context.Background())
		quitter := func(k *terminalapi.Keyboard) {
			if k.Key == 'q' || k.Key == 'Q' {
				cancel()
			}
		}
		if err := termdash.Run(ctx, terminalLayer, containerLayer, termdash.KeyboardSubscriber(quitter)); err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(dashboardCmd)
}

func rollText(rollingText *text.Text, value <-chan string) {
	for {
		select {
		case v := <-value:
			err := rollingText.Write(v)
			if err != nil {
				panic(err)
			}
		default:
			// no value on the channel, continue without blocking
		}
	}
}
