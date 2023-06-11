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
		leftContainer := container.Left(container.Border(linestyle.Light))
		rightContainer := container.Right(
			container.SplitHorizontal(
				container.Top(
					container.Border(linestyle.Light),
				),
				container.Bottom(
					container.SplitVertical(
						container.Left(
							container.Border(linestyle.Light),
						),
						container.Right(
							container.Border(linestyle.Light),
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
