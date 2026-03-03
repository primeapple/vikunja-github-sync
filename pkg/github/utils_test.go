package github

import "testing"

func TestGetNextPageUrl(t *testing.T) {
	tests := []struct {
		name     string
		link     string
		expected string
	}{
		{
			name:     "empty header",
			link:     "",
			expected: "",
		},
		{
			name:     "valid link header with next",
			link:     `<https://api.github.com/issues?page=2>; rel="next"`,
			expected: "https://api.github.com/issues?page=2",
		},
		{
			name:     "link header without rel=next",
			link:     `<https://api.github.com/issues>; rel="prev"`,
			expected: "",
		},
		{
			name:     "multiple links, extract next",
			link:     `<https://api.github.com/issues>; rel="prev", <https://api.github.com/issues?page=2>; rel="next", <https://api.github.com/issues?page=3>; rel="last"`,
			expected: "https://api.github.com/issues?page=2",
		},
		{
			name:     "next link with spaces",
			link:     ` <https://api.github.com/issues?page=2> ; rel="next" `,
			expected: "https://api.github.com/issues?page=2",
		},
		{
			name:     "multiple links with whitespace",
			link:     `<https://api.github.com/issues?page=1>; rel="first", <https://api.github.com/issues?page=2>; rel="next"`,
			expected: "https://api.github.com/issues?page=2",
		},
		{
			name:     "no link header present",
			link:     "text without links",
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getNextPageUrl(tt.link)
			if result != tt.expected {
				t.Errorf("getNextPageUrl() = %q, want %q", result, tt.expected)
			}
		})
	}
}
