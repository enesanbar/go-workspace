package maximum_path_sum

func FindMaxPathInTriangle(input [][]int) int {
	length := len(input)

	// initialize cost array
	cost := make([][]int, length)
	rows := make([]int, length*length)
	for i := 0; i < length; i++ {
		cost[i] = rows[i*length : (i+1)*length]
	}

	// get cost of first row
	cost[0][0] = input[0][0]

	// cover left edge
	for i := 1; i < length; i++ {
		cost[i][0] = cost[i-1][0] + input[i][0]
	}

	// cover rest
	for i := 1; i < length; i++ {
		for j := 1; j <= i; j++ {
			cost[i][j] = Max(cost[i-1][j-1], cost[i-1][j]) + input[i][j]
		}
	}

	return Max(cost[length-1]...)
}

func Max(input ...int) int {
	var max int
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	return max
}
