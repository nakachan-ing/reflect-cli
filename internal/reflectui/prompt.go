package reflectui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nakachan-ing/reflect-cli/model"
)

func RunInteractivePrompt(template *model.ReflectTemplate) ([]string, error) {
	var responses []string

	fmt.Printf("🧠 Starting Reflect Prompt: %s\n\n", template.Title)

	reader := bufio.NewReader(os.Stdin)

	for i, question := range template.Prompts {
		fmt.Printf("Q%d. %s\n", i+1, question)
		fmt.Println("（複数行を入力してください。完了したら Ctrl+D（または Ctrl+Z）で終了）")

		var lines []string
		for {
			fmt.Print("> ")
			line, err := reader.ReadString('\n')
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				return nil, fmt.Errorf("failed to read input: %w", err)
			}
			lines = append(lines, strings.TrimRight(line, "\r\n"))
		}

		response := strings.Join(lines, "\n")
		responses = append(responses, response)
		fmt.Println()
	}

	return responses, nil
}
