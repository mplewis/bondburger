package split

import (
	"regexp"
	"strings"
)

func SplitSentences(all string) []string {
	all = strings.TrimSpace(all)
	results := regexp.MustCompile(`([^\s]+)\.\s+`).FindAllStringIndex(all, -1)
	sentences := make([]string, len(results))
	lastPos := 0
	for i := 0; i < len(results); i++ {
		currPos := results[i][1]
		sentences[i] = all[lastPos:currPos]
		lastPos = currPos
	}
	sentences = append(sentences, all[lastPos:])
	for i, s := range sentences {
		sentences[i] = strings.TrimSpace(s)
	}
	return sentences
}
