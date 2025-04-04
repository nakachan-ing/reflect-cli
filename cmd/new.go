/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:     "new",
	Short:   "", // あとで説明は追記
	Aliases: []string{"n"},
}

var newFleetingCmd = &cobra.Command{
	Use:   "fleeting [title]",
	Short: "Add a new fleeting note",
	// Args:    cobra.ExactArgs(1), 今後の引数による
	Aliases: []string{"f"},
	Run: func(cmd *cobra.Command, args []string) {
		// ここにnew fleeting の処理を実装
	},
}

func init() {
	newCmd.AddCommand(newFleetingCmd)
	rootCmd.AddCommand(newCmd)
}
