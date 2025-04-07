package model

import (
	"fmt"
	"regexp"
	"strings"
)

// 最低限のバリデーションを実装
// 禁則文字が存在する場合、そのタグ作成をスキップ
func ValidateTags(tags []string) ([]string, error) {
	var valid []string
	var invalid []string
	var tagRegexp = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}
		if !tagRegexp.MatchString(tag) {
			invalid = append(invalid, tag)
			continue
		}
		if !contains(valid, tag) {
			valid = append(valid, tag)
		}
	}

	if len(invalid) > 0 {
		fmt.Printf("Warning: Some tags were invalid and skipped: %v\n", invalid)
	}

	// 今は、常にerrorでnilを返すが、今後、error実装検討
	return valid, nil
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
