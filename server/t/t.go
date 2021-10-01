package t

import (
	"regexp"
	"strings"
)

type Film struct {
	Title string   `json:"title"`
	Year  int      `json:"year"`
	Actor string   `json:"actor"`
	Plot  []string `json:"plot"`
}

func (f *Film) Slug() (slug string) {
	slug = strings.ToLower(f.Title)
	slug = regexp.MustCompile(`[^\w\s]`).ReplaceAllLiteralString(slug, "")
	slug = regexp.MustCompile(`\s+`).ReplaceAllLiteralString(slug, "_")
	return
}
