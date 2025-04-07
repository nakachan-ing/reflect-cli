/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nakachan-ing/reflect-cli/config"
	"github.com/nakachan-ing/reflect-cli/internal/noteio"
	"github.com/nakachan-ing/reflect-cli/internal/templateio"
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
		// ここでconfigを読み込む
		config, err := config.LoadConfig()
		if err != nil {
			log.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		// ここにnew fleeting の処理を実装
		validatedSubType, err := model.IsSubType(subType)
		if err != nil {
			log.Printf("Error: %v\n", err)
			os.Exit(2)
		}
		// fmt.Println(validatedSubType)

		// ここにtitleのバリデーションを実装
		slug, err := utils.Slugify(title)
		if err != nil {
			// slugがなくてもファイル作成できるように Warningのみにしておく
			log.Printf("Warning: %v\n", err)
		}
		// fmt.Println(slug)

		// ここにTagのバリデーションを実装
		// fmt.Println(tags)
		validatedTags, err := model.ValidateTags(tags)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
		// fmt.Println(validatedTags)

		// ここにsourceのバリデーションを実装
		// --typeがliteratureの場合は、必須項目
		// それ以外の場合は、任意項目
		if err = model.IsSourceSpecified(validatedSubType, source); err != nil {
			log.Printf("Error: %v\n", err)
			os.Exit(2)
		}
		// fmt.Println(source)

		// ここにissueのバリデーションを実装
		validatedIssue, warning := utils.ValidateIssueURL(issue)
		if warning != "" {
			log.Printf("Warning: %v\n", warning)
		}
		fmt.Println(validatedIssue)

		newTags := model.MapTags(validatedTags)
		newNote := model.MapNote(title, subType, slug, source, validatedIssue, newTags)

		// debug
		// fmt.Println(newNote)
		// fmt.Println("ID:", newNote.ID)
		// fmt.Println("NoteType:", newNote.NoteType)
		// fmt.Println("SubType:", newNote.SubType)
		// fmt.Println("CreatedAt:", newNote.CreatedAt)
		// fmt.Println("UpdatedAt:", newNote.UpdatedAt)
		// fmt.Println("Archived:", newNote.Archived)
		// fmt.Println("Deleted:", newNote.Deleted)
		// fmt.Println("Reflected:", newNote.Reflected)
		// fmt.Println("FilePath:", newNote.FilePath)
		// fmt.Println("Slug:", newNote.Slug)
		// fmt.Println("Source:", newNote.Source)
		// fmt.Println("LinkedIssue:", newNote.LinkedIssue)
		// fmt.Println("LinkedNotes:", newNote.LinkedNotes)
		// fmt.Println("Tags:")
		// for _, tag := range newNote.Tags {
		// 	fmt.Printf("  %v\n", tag.Name)
		// }

		frontMatterBytes, err := model.MapFrontMatter(title, subType, source, validatedIssue, validatedTags)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
		// fmt.Println(string(frontMatterBytes))

		// fileに書き出し
		fleetingTempletePath := filepath.Join(config.TemplateDir, "fleeting", config.Language, fmt.Sprintf("%s.md", newNote.SubType))
		templateContent, err := templateio.LoadFleetingTemplate(fleetingTempletePath)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}

		filePath, err := noteio.WriteNoteFile(newNote, string(frontMatterBytes), templateContent, *config)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}

		fmt.Printf("✅ Created new fleeting note: %s\n", filepath.Base(filePath))
		fmt.Printf("📁 Path: %s\n", filePath)

		if len(newNote.Tags) > 0 {
			var tagNames []string
			for _, t := range newNote.Tags {
				tagNames = append(tagNames, t.Name)
			}
			fmt.Printf("🏷️ Tags: %s\n", strings.Join(tagNames, ", "))
		}

		if newNote.LinkedIssue != "" {
			fmt.Printf("🔗 Linked Issue: %s\n", newNote.LinkedIssue)
		}
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
