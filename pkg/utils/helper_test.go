package utils

import (
	"strconv"
	"testing"
)

func TestPop(t *testing.T) {
	t.Run("should give first element and remove it", func(t *testing.T) {
		list := []int{1, 2}
		got, ok := Pop(&list)
		want := 1
		AssertTrue(t, ok)
		AssertEquals(t, got, want)
		AssertDeepEquals(t, list, []int{2})
	})

	t.Run("should return zero value if list is empty", func(t *testing.T) {
		list := []int{}
		got, ok := Pop(&list)
		want := 0
		AssertFalse(t, ok)
		AssertEquals(t, got, want)
		AssertDeepEquals(t, list, []int{})
	})
}

func TestFilter(t *testing.T) {
	greaterThanFive := func(a int) bool {
		return a > 5
	}

	t.Run("should filter properly", func(t *testing.T) {
		got := Filter([]int{4, 6}, greaterThanFive)
		want := []int{6}
		AssertDeepEquals(t, got, want)
	})

	t.Run("should work with empty list", func(t *testing.T) {
		got := Filter([]int{}, greaterThanFive)
		want := []int{}
		AssertDeepEquals(t, got, want)
	})
}

func TestMap(t *testing.T) {
	toString := func(a int) string {
		return strconv.Itoa(a)
	}

	t.Run("should map properly", func(t *testing.T) {
		got := Map([]int{4, 6}, toString)
		want := []string{"4", "6"}
		AssertDeepEquals(t, got, want)
	})

	t.Run("should work with empty list", func(t *testing.T) {
		got := Map([]int{}, toString)
		want := []string{}
		AssertDeepEquals(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	sum := func(a, b int) int {
		return a + b
	}

	t.Run("should use the starting accumulator", func(t *testing.T) {
		got := Reduce([]int{1, 2}, 1000, sum)
		want := 1003
		AssertEquals(t, got, want)
	})

	t.Run("should work with empty list", func(t *testing.T) {
		got := Reduce([]int{}, 1000, sum)
		want := 1000
		AssertEquals(t, got, want)
	})
}
