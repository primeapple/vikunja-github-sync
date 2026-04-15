package github

import (
	"fmt"
)

func NewGithubIssue(repository string, issueNumber int) *GithubIssue {
	return &GithubIssue{
		repository: repository,
		issueNumber: issueNumber,
	}
}

type GithubIssue struct {
	repository string
	issueNumber int
}

func (issue *GithubIssue) Id() string {
	return fmt.Sprintf("%s#%d", issue.repository, issue.issueNumber)
}
