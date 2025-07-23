package main

import (
	"fmt"

	"github.com/primeapple/vikunja-github-sync/pkg/vikunja"
)

func main() {
	vikunja := vikunja.NewClient()
	default_project_id := vikunja.GetDefaultProjectId()
	fmt.Println(*default_project_id)
}

// func get_vikunja_projects() {
// 	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/projects", URL), nil)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", TOKEN))
// 	response, err := http.DefaultClient.Do(req)
//
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()
//
// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	//Convert the body to type string
// 	sb := string(body)
// 	fmt.Printf(sb)
// }
//
// func create_vikunja_task() {
// 	data := []byte(`{"key": "value"}`)
// 	http.NewRequest(http.MethodPut, "https://try.vikunja.io/api/v1/projects/{id}/tasks")
// }
