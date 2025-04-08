package model

type Config struct {
	BaseDir        string `yaml:"baseDir"`
	TemplateDir    string `yaml:"templateDir"`
	Language       string `yaml:"language"`
	ZettelJsonPath string `yaml:"zettelJsonPath"`
}
