package sync

import (
	"fmt"

	"github.com/primeapple/vikunja-github-sync/pkg/github"
)

func Sync() error {
	// vikunja := vikunja.NewClient()
	// default_project_id, err := vikunja.GetDefaultProjectId()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("Default project id: %d\n", *default_project_id)
	//
	// err = vikunja.CreateTask(*default_project_id, "Created via API", "Created by vikunja-github-sync")
	// if err != nil {
	// 	return err
	// }
	github := github.NewClient()
	issue_urls, err := github.GetAssignedOpenIssues()
	if err != nil {
		return err
	}
	fmt.Printf("Default project id: %v", issue_urls)

	return nil
}
