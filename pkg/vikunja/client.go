package vikunja

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type VikunjaClient struct {
	baseUrl string
	token   string
}

func NewClient() *VikunjaClient {
	baseUrl := os.Getenv("VIKUNJA_URL")
	token := os.Getenv("VIKUNJA_TOKEN")

	return &VikunjaClient{
		baseUrl: baseUrl,
		token:   token,
	}
}

type UserResponse struct {
	Settings struct {
		DefaultProjectId int `json:"default_project_id"`
	} `json:"settings"`
}

func (client VikunjaClient) GetDefaultProjectId() (*int, error) {
	body, err := client.request(http.MethodGet, "user", nil)
	if err != nil {
		return nil, err
	}

	var resp UserResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	defaultProjectId := resp.Settings.DefaultProjectId

	// If it is 0 we don't have a default project set yet
	if defaultProjectId == 0 {
		return nil, fmt.Errorf("No default project found")
	}

	return &defaultProjectId, nil
}

type PutTaskRequest struct {
	Title string `json:"title"`
}

func (client VikunjaClient) CreateTask(projectId int, title string) error {
	reqBody := PutTaskRequest{Title: title}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	_, err = client.request(http.MethodPut, fmt.Sprintf("projects/%d/tasks", projectId), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	return nil
}

func (client VikunjaClient) request(method string, path string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", client.baseUrl, path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.token))
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
