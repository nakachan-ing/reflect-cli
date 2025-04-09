package templateio

import (
	"fmt"
	"strings"

	"github.com/nakachan-ing/reflect-cli/model"
)

func BuildReflectBody(template *model.ReflectTemplate, responses []string) string {
	var sb strings.Builder

	for i, section := range template.AbstractDimensions {
		sb.WriteString(fmt.Sprintf("## %s\n", section))
		sb.WriteString(responses[i])
		sb.WriteString("\n\n")
	}

	return sb.String()
}
