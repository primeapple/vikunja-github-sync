package task

import (
	"slices"

	"github.com/primeapple/vikunja-github-sync/pkg/utils"
)

func Minus(tasks, others Tasks) Tasks {
	return utils.Filter(tasks, func(t Task) bool {
		return !slices.Contains(others, t)
	})
}
