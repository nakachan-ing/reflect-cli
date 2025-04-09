package templateio

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nakachan-ing/reflect-cli/model"
	"gopkg.in/yaml.v3"
)

func LoadReflectTemplate(subtype, lang string, config model.Config) (*model.ReflectTemplate, error) {
	path := filepath.Join(config.TemplateDir, "reflect", lang, fmt.Sprintf("%s.yaml", subtype))
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tmpl model.ReflectTemplate
	if err := yaml.Unmarshal(data, &tmpl); err != nil {
		return nil, err
	}
	return &tmpl, nil
}
