/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/nakachan-ing/reflect-cli/config"
	"github.com/nakachan-ing/reflect-cli/internal/noteio"
	"github.com/nakachan-ing/reflect-cli/internal/reflectui"
	"github.com/nakachan-ing/reflect-cli/internal/templateio"
	"github.com/nakachan-ing/reflect-cli/model"
	"github.com/spf13/cobra"
)

var interactive bool
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
			fmt.Println("Reflect Responses:")
			for i, q := range template.Prompts {
				fmt.Printf("Q%d: %s\n", i+1, q)
				fmt.Printf("A%d: %s\n\n", i+1, responses[i])
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
	reflectCmd.Flags().StringVar(&reflectType, "type", "", "Specify note subtype (idea, question, ...)")
	reflectCmd.Flags().StringVar(&reflectLanguage, "lang", "", "Specify language for prompts (e.g. ja, en)")

}
