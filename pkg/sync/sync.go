package sync

import (
	"fmt"

	"github.com/primeapple/vikunja-github-sync/pkg/vikunja"
)

func Sync() error {
	vikunja := vikunja.NewClient()
	default_project_id, err := vikunja.GetDefaultProjectId()
	if err != nil {
		return err
	}
	fmt.Printf("Default project id: %d\n", *default_project_id)

	err = vikunja.CreateTask(*default_project_id, "Created via API")
	if err != nil {
		return err
	}

	return nil
}
