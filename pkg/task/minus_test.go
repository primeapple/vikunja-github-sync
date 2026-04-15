package task

import (
	"testing"

	"github.com/primeapple/vikunja-github-sync/pkg/github"
	"github.com/primeapple/vikunja-github-sync/pkg/utils"
)

func testIssue(number int) github.GithubIssue {
	return github.NewGithubIssue("primeapple/vikunja-github-sync", number)
}

func TestMinus(t *testing.T) {
	tasks := []Task{testIssue(1), testIssue(2)}

	t.Run("should return empty tasks if both are the same", func(t *testing.T) {
		want := []Task{}
		got := Minus(tasks, tasks)
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should return empty tasks if both are empty", func(t *testing.T) {
		want := []Task{}
		got := Minus([]Task{}, []Task{})
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should return all tasks if others is empty", func(t *testing.T) {
		want := tasks
		got := Minus(tasks, []Task{})
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should return empty tasks if tasks is empty", func(t *testing.T) {
		want := []Task{}
		got := Minus([]Task{}, tasks)
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should return only tasks not in others", func(t *testing.T) {
		want := []Task{issue2}
		got := Minus(tasks, []Task{issue1})
		utils.AssertDeepEquals(t, got, want)
	})

	t.Run("should not be affected by elements in others that are not in tasks", func(t *testing.T) {
		want := tasks
		got := Minus(tasks, []Task{issue3})
		utils.AssertDeepEquals(t, got, want)
	})
}
