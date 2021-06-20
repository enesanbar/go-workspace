package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	maximumPathSum "github.com/enesanbar/workspace/golang/projecteuler/18-maximum-path-sum"
)

func main() {
	input := ReadTriangularInputFromStdin()
	triangle := maximumPathSum.FindMaxPathInTriangle(input)
	fmt.Println(triangle)
}

func ReadTriangularInputFromStdin() [][]int {
	input := make([][]int, 0)
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		text := reader.Text()
		split := strings.Split(text, " ")
		line := make([]int, 0)
		for _, v := range split {
			intValue, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			line = append(line, intValue)
		}
		input = append(input, line)
	}

	return input
}
