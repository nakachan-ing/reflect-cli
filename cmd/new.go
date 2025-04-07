/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/nakachan-ing/reflect-cli/model"
	"github.com/nakachan-ing/reflect-cli/utils"
	"github.com/spf13/cobra"
)

// Argument variables
var subType string
var title string
var tags []string
var source string
var issue string

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
		validatedSubType, err := model.IsSubType(subType)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(2)
		}
		fmt.Println(validatedSubType)

		// ここにtitleのバリデーションを実装
		slug, err := utils.Slugify(title)
		if err != nil {
			// slugがなくてもファイル作成できるように Warningのみにしておく
			fmt.Println("Warning:", err)
		}
		fmt.Println(slug)

		// ここにTagのバリデーションを実装
		// fmt.Println(tags)
		validatedTags, err := model.ValidateTags(tags)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(validatedTags)

		// ここにsourceのバリデーションを実装
		// --typeがliteratureの場合は、必須項目
		// それ以外の場合は、任意項目
		if err = model.IsSourceSpecified(validatedSubType, source); err != nil {
			fmt.Println("Error:", err)
			os.Exit(2)
		}
		fmt.Println(source)

		// ここにissueのバリデーションを実装
		validatedIssue, warning := utils.ValidateIssueURL(issue)
		if warning != "" {
			fmt.Println("Warning:", warning)
		}
		fmt.Println(validatedIssue)

	},
}

func init() {
	newCmd.AddCommand(newFleetingCmd)
	rootCmd.AddCommand(newCmd)

	// Options
	newFleetingCmd.Flags().StringVarP(&subType, "type", "t", "", "Specify fleeting note type")
	newFleetingCmd.MarkFlagRequired("type")
	newFleetingCmd.Flags().StringVar(&title, "title", "", "Specify fleeting note title")
	newFleetingCmd.Flags().StringSliceVar(&tags, "tags", []string{}, "Specify tags")
	newFleetingCmd.Flags().StringVarP(&source, "source", "s", "", "Specify literature source")
	newFleetingCmd.Flags().StringVarP(&issue, "issue", "i", "", "Specify related issue")
}
