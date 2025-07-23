package vikunja

import (
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

func (client VikunjaClient) GetDefaultProjectId() *int {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user", client.baseUrl), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.token))
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var resp UserResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		panic(err)
	}

	defaultProjectId := resp.Settings.DefaultProjectId

	// If it is 0 we don't have a default project set yet
	if defaultProjectId == 0 {
		return nil
	}

	return &defaultProjectId
}
