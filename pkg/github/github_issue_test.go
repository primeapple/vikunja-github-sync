package github

import (
	"testing"

	"github.com/primeapple/vikunja-github-sync/pkg/utils"
)

func TestId(t *testing.T) {
	t.Run("should return correct id", func(t *testing.T) {
		issue := NewGithubIssue("primeapple/vikunja-github-sync", 1)
		utils.AssertEquals(t, issue.Id(), "primeapple/vikunja-github-sync#1")
	})
}
