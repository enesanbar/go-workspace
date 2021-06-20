package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	namesScores "github.com/enesanbar/workspace/golang/projecteuler/22-names-scores"
)

func main() {
	f, err := os.Open("./p022_names.txt")
	if err != nil {
		panic(err)
	}

	names := ReadNames(f)
	sort.Strings(names)
	score := namesScores.CalculateNameScore(names)
	fmt.Println("Score:", score)
}

func ReadNames(reader io.Reader) []string {
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	split := strings.Split(string(all), ",")
	result := make([]string, 0, len(split))
	for _, name := range split {
		runes := []rune(name)
		result = append(result, string(runes[1:len(runes)-1]))
	}

	return result
}
