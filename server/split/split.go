package split

import (
	"regexp"
	"strings"
)

// https://github.com/mediacloud/sentence-splitter/blob/9af995f451a193898953dfe3926d1beefb45afb6/sentence_splitter/non_breaking_prefixes/en.txt
var _abbrs = []string{
	"a", "adj", "adm", "adv", "al", "apr", "art", "asst", "aug", "b", "bart", "bldg", "brig", "bros", "c", "capt", "cf",
	"cmdr", "co", "col", "comdr", "con", "corp", "cpl", "d", "dec", "dr", "drs", "e", "e.g", "ens", "esp", "etc", "f",
	"feb", "fig", "g", "gen", "gov", "h", "hon", "hosp", "hr", "i", "i.e", "inc", "insp", "j", "jan", "jr", "jul", "jun",
	"k", "l", "lt", "m", "maj", "mar", "messrs", "mlle", "mm", "mme", "mr", "mrs", "ms", "msgr", "n", "no", "nos", "nov",
	"nr", "o", "oct", "okt", "op", "ord", "p", "pfc", "ph", "ph.d", "phd", "pp", "prof", "pvt", "q", "r", "rep", "reps",
	"res", "rev", "rt", "s", "sen", "sens", "sep", "sept", "sfc", "sgt", "sr", "st", "supt", "surg", "t", "u",
	"u.k", "u.s", "v", "vs", "w", "x", "y", "z",
}
var abbrs = map[string]struct{}{}

func init() {
	for _, abbr := range _abbrs {
		abbrs[abbr+"."] = struct{}{}
	}
}

func Sentences(all string) []string {
	all = strings.TrimSpace(all)
	results := regexp.MustCompile(`([^\s]+)\.\s+`).FindAllStringIndex(all, -1)
	sentences := make([]string, len(results))
	lastPos := 0
nextWord:
	for i := 0; i < len(results); i++ {
		lastRaw := all[results[i][0]:results[i][1]]
		lastWord := strings.ToLower(strings.TrimSpace(lastRaw))
		if _, breakIsAbbr := abbrs[lastWord]; breakIsAbbr {
			continue nextWord
		}

		currPos := results[i][1]
		sentences[i] = all[lastPos:currPos]
		lastPos = currPos
	}
	sentences = append(sentences, all[lastPos:])

	allTrimmed := []string{}
	for _, s := range sentences {
		trimmed := strings.TrimSpace(s)
		if trimmed == "" {
			continue
		}
		allTrimmed = append(allTrimmed, trimmed)
	}
	return allTrimmed
}
