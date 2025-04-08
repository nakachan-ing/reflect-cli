package utils

import (
	"errors"
	"log"
	"os"

	"github.com/nakachan-ing/reflect-cli/internal/store/jsonstore"
)

func HandleZettelJsonError(err error) {
	var notExistErr *jsonstore.ZettelJsonNotExistError
	var parseErr *jsonstore.ZettelJsonParseError
	var readErr *jsonstore.ZettelJsonReadError

	switch {
	case errors.As(err, &notExistErr):
		log.Printf("❌ Error: %s", notExistErr)
		os.Exit(1)
	case errors.As(err, &parseErr):
		log.Printf("❌ Error: %s", parseErr)
		os.Exit(1)
	case errors.As(err, &readErr):
		log.Printf("❌ Error: %s", readErr)
		os.Exit(1)
	default:
		log.Printf("❌ Unknown Error: %v", err)
		os.Exit(1)
	}
}
