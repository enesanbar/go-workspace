package main

import (
	"fmt"
	"sort"
)

func main() {
	ways := getWays(3, []int64{8, 3, 1, 2})
	fmt.Println(ways)
}

func getWays(n int32, c []int64) int64 {
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
	c = append([]int64{0}, c...)

	// initialize cost array
	rowCount := len(c)
	colCount := n + 1
	cost := make([][]int64, rowCount)
	rows := make([]int64, rowCount*int(colCount))
	for i := 0; i < rowCount; i++ {
		cost[i] = rows[i*int(colCount) : (i+1)*int(colCount)]
	}

	// base condition
	cost[0][0] = 1
	for i := 0; i < int(colCount); i++ {
		cost[0][i] = 0
	}

	for i := 0; i < rowCount; i++ {
		cost[i][0] = 1
	}

	// go over cells
	for i := 1; i < rowCount; i++ {
		for j := 1; j < int(colCount); j++ {
			currentTargetAmount := j
			currentCoin := c[i]
			neededAmount := int64(currentTargetAmount) - currentCoin

			costWithoutCurrentCoin := cost[i-1][j]
			cost[i][j] = costWithoutCurrentCoin
			if neededAmount >= 0 {
				cost[i][j] += cost[i][neededAmount]
			}
		}
	}

	return cost[rowCount-1][colCount-1]
}

func PrintMatrix(matrix [][]int64) {
	for _, row := range matrix {
		for _, cell := range row {
			fmt.Printf("%d ", cell)
		}
		fmt.Println()
	}
}
