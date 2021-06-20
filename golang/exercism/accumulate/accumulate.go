// Package accumulate provides helper functions to convert list of strings
package accumulate

// Accumulate returns the list of strings with the converter function applied to each item in the list
func Accumulate(words []string, converter func(string) string) []string {
	result := make([]string, len(words))

	for i, item := range words {
		result[i] = converter(item)
	}

	return result
}
