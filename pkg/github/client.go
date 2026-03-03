package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type GithubClient struct {
	token string
}

func NewClient() *GithubClient {
	token := os.Getenv("GITHUB_TOKEN")

	return &GithubClient{
		token: token,
	}
}

type GetIssuesResponse struct {
	Number     int `json:"number"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
}

// see https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#list-issues-assigned-to-the-authenticated-user
func (client GithubClient) GetAssignedOpenIssues() ([]GetIssuesResponse, error) {
	var allIssues []GetIssuesResponse
	url := "https://api.github.com/issues?state=open&filter=assigned&per_page=100"

	for url != "" {
		body, headers, err := client.request(http.MethodGet, url, nil)
		if err != nil {
			return nil, fmt.Errorf("Can't fetch assigned github issues: %w", err)
		}

		var resp []GetIssuesResponse
		if err := json.Unmarshal(body, &resp); err != nil {
			return nil, fmt.Errorf("Can't extract assigned github issues: %w", err)
		}
		allIssues = append(allIssues, resp...)

		url = getNextPageUrl(headers.Get("Link"))
		println(url)
	}

	return allIssues, nil
}

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

func (client GithubClient) request(method string, url string, body io.Reader) ([]byte, http.Header, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.token))
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	if method != http.MethodGet {
		req.Header.Add("Content-Type", "application/json")
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode >= 400 {
		return nil, nil, fmt.Errorf("Error in response, status %d, body %s, %v", response.StatusCode, string(respBody), *response)
	}

	return respBody, response.Header, nil
}
