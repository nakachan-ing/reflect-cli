package model

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type Tag struct {
	ID        string     `json:"id"` // 例: t-001
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	Deleted   bool       `json:"deleted"`
}

func GetNextTagID(tags []Tag) string {
	maxSeqID := 0
	re := regexp.MustCompile(`t(\d+)`) // "tXXX" の数字部分を抽出する正規表現

	// SeqID の最大値を探す
	for _, tag := range tags {
		match := re.FindStringSubmatch(tag.ID)
		if match != nil {
			seq, err := strconv.Atoi(match[1]) // "XXX" 部分を整数に変換
			if err == nil && seq > maxSeqID {
				maxSeqID = seq
			}
		}
	}

	// 新しいIDを生成
	newSeqID := maxSeqID + 1

	// 999 までは3桁ゼロ埋め、それ以上はそのまま
	if newSeqID < 1000 {
		return fmt.Sprintf("t%03d", newSeqID) // 3桁ゼロ埋め
	}
	return fmt.Sprintf("t%d", newSeqID) // 1000以上はゼロ埋めなし
}
