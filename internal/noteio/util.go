package noteio

import (
	"fmt"
	"regexp"
)

func ExtractNoteID(filename string) (string, error) {
	re := regexp.MustCompile(`^(\d{8}T\d{6})`)
	match := re.FindStringSubmatch(filename)
	if len(match) >= 2 {
		return match[1], nil
	}
	return "", fmt.Errorf("note ID not found in filename: %s", filename)
}
