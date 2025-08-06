package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

type GetUserResponse struct {
	Login string `json:"login"`
}

func (client GithubClient) getUsername() (string, error) {
	body, err := client.request(http.MethodGet, "user", nil)
	if err != nil {
		return "", err
	}

	var resp GetUserResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", err
	}

	return resp.Login, nil
}

type SearchIssuesResponse struct {
	Items []struct {
		Url string `json:"url"`
	} `json:"items"`
}

func (client GithubClient) GetAssignedOpenIssues() ([]string, error) {
	username, err := client.getUsername()
	if err != nil {
		return nil, err
	}

	body, err := client.request(http.MethodGet, fmt.Sprintf("search/issues?q=is:open+is:issue+assignee:%s&per_page=100", username), nil)
	if err != nil {
		return nil, err
	}

	var resp SearchIssuesResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	var assigned_issues_urls []string
	for _, issue := range resp.Items {
		assigned_issues_urls = append(assigned_issues_urls, issue.Url)
	}

	return assigned_issues_urls, nil
}

func (client GithubClient) request(method string, path string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("https://api.github.com/%s", path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.token))
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	if method != http.MethodGet {
		req.Header.Add("Content-Type", "application/json")
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("Error in response, status %d, body %s, %v", response.StatusCode, string(respBody), *response)
	}

	return respBody, nil
}
