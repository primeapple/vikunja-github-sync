package github

import "strings"

func getNextPageUrl(linkHeader string) string {
	if linkHeader == "" {
		return ""
	}

	links := strings.SplitSeq(linkHeader, ",")
	for link := range links {
		parts := strings.Split(strings.TrimSpace(link), ";")
		if len(parts) == 2 && strings.Contains(parts[1], `rel="next"`) {
			url := strings.Trim(strings.TrimSpace(parts[0]), "<>")
			return url
		}
	}
	return ""
}
