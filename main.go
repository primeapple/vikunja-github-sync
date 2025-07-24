package main

import (
	"fmt"
	"os"

	"github.com/primeapple/vikunja-github-sync/pkg/sync"
)

func main() {
	err := sync.Sync()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
