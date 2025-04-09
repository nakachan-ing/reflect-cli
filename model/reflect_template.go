package model

type ReflectTemplate struct {
	Type               string   `yaml:"type"`
	Title              string   `yaml:"title"`
	AbstractDimensions []string `yaml:"abstract_dimensions"`
	Prompts            []string `yaml:"prompts"`
}
