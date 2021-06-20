package main

import (
	"log"
	"os"

	_ "github.com/enesanbar/workspace/projects/search-feed/matchers"
	"github.com/enesanbar/workspace/projects/search-feed/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
