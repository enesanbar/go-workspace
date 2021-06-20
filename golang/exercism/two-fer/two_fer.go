// Package twofer implements a utility function to returns "two for one" string with or without name.
package twofer

import "fmt"

// ShareWith returns the string "One for X, one for me." where x is either "you" or the parameter 'name'
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
