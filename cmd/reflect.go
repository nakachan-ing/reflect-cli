/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/nakachan-ing/reflect-cli/config"
	"github.com/nakachan-ing/reflect-cli/internal/noteio"
	"github.com/nakachan-ing/reflect-cli/internal/reflectui"
	"github.com/nakachan-ing/reflect-cli/internal/store/jsonstore"
	"github.com/nakachan-ing/reflect-cli/internal/templateio"
	"github.com/nakachan-ing/reflect-cli/model"
	"github.com/nakachan-ing/reflect-cli/utils"
	"github.com/spf13/cobra"
)

var interactive bool
var reflectTitle string
var reflectType string
var reflectLanguage string

// reflectCmd represents the reflect command
var reflectCmd = &cobra.Command{
	Use:   "reflect [notePath]",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// ここでconfigを読み込む
		config, err := config.LoadConfig()
		if err != nil {
			log.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		filePath := args[0]

		frontMatter, _, err := noteio.ParseNoteFile[model.FrontMatter](filePath)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}

		if subType == "" {
			// frontmatter から subtype を読み取る
			subType = frontMatter.SubType
		}

		if language == "" {
			language = config.Language
		}

		template, err := templateio.LoadReflectTemplate(subType, language, *config)
		if err != nil {
			log.Printf("Error: %v\n", err)

		}

		if interactive {
			responses, err := reflectui.RunInteractivePrompt(template)
			if err != nil {
				log.Printf("Error: %v\n", err)
			}

			// ここにtitleのバリデーションを実装
			slug, err := utils.Slugify(reflectTitle)
			if err != nil {
				// slugがなくてもファイル作成できるように Warningのみにしておく
				log.Printf("Warning: %v\n", err)
			}

			existNotes, err := jsonstore.LoadNotes(*config)
			if err != nil {
				log.Printf("Error: %s", err)
				os.Exit(1)
			}

			id, err := noteio.ExtractNoteID(filepath.Base(filePath))
			if err != nil {
				log.Printf("Error: %s", err)
			}

			var fleetingNotes []*model.Note
			found := false
			for i := range existNotes {
				if existNotes[i].ID == id {
					fleetingNotes = append(fleetingNotes, &existNotes[i])
					found = true
				}
			}

			if !found {
				log.Printf("Note with ID %s not found", id)
			}

			existTags, err := jsonstore.LoadTags(*config)
			if err != nil {
				log.Printf("Error: %s", err)
			}

			var permTags []*model.Tag
			for _, tag := range frontMatter.Tags {
				for i := range existTags {
					if existTags[i].Name == tag {
						permTags = append(permTags, &existTags[i])
					}
				}
			}

			perm := model.MapReflectToPermanent(reflectTitle, template.Type, slug, frontMatter.Source, frontMatter.LinkedIssue, fleetingNotes, permTags, responses, config)

			frontMatterBytes, err := model.MapFrontMatter(perm.Title, "permanent", string(perm.SubType), frontMatter.Source, frontMatter.LinkedIssue, frontMatter.Tags, filepath.Base(filePath))
			if err != nil {
				log.Printf("Error: %v\n", err)
			}

			body := templateio.BuildReflectBody(template, responses)

			filePath, err := noteio.WriteNoteFile(perm, string(frontMatterBytes), body, *config)
			if err != nil {
				log.Printf("Error: %v\n", err)
			}

			fmt.Printf("\n🧠 Reflect completed! Review below:\n\n")

			// 回答確認
			for i, q := range template.Prompts {
				fmt.Printf("Q%d: %s\nA%d: %s\n\n", i+1, q, i+1, responses[i])
			}

			// 保存完了通知
			fmt.Printf("✅ Permanent note saved: %s\n", filepath.Base(filePath))
			fmt.Printf("📁 Path: %s\n", filePath)

			if err = jsonstore.InsertNoteToJson(perm, config); err != nil {
				utils.HandleZettelJsonError(err)
			}

		} else {
			// 今はバッチ未対応、警告だけでOK
			fmt.Println("Non-interactive mode is not supported yet. Use --interactive.")
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(reflectCmd)

	// options
	reflectCmd.Flags().BoolVar(&interactive, "interactive", true, "Run in interactive prompt mode")
	reflectCmd.Flags().StringVar(&reflectTitle, "title", "", "Specify title for the permanent note")
	reflectCmd.MarkFlagRequired("title")
	reflectCmd.Flags().StringVar(&reflectType, "type", "", "Specify note subtype (idea, question, ...)")
	reflectCmd.Flags().StringVar(&reflectLanguage, "lang", "", "Specify language for prompts (e.g. ja, en)")

}
