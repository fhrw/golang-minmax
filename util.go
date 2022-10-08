package golangMinMax

func isTerminal(state [][]int) bool {
	// if any row has three of the same
	for _, v := range state {
		if sum(v) == 3 || sum(v) == 6 {
			return true
		}
	}
	// if any col has three
	for i := range state {
		col := getCol(state, i)
		if sum(col) == 3 || sum(col) == 6 {
			return true
		}
	}
	// if either diag has three
	diagOne := []int{state[0][0], state[1][1], state[2][2]}
	diagTwo := []int{state[0][2], state[1][1], state[2][0]}
	if sum(diagOne) == 3 || sum(diagTwo) == 3 || sum(diagOne) == 6 || sum(diagTwo) == 6 {
		return true
	}
	// draw
	if len(availSquares(state)) == 0 {
		return true
	}
	// not terminal
	return false
}

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func getCol(s [][]int, idx int) []int {
	col := []int{}
	for _, r := range s {
		col = append(col, r[idx])
	}
	return col
}

func availSquares(s [][]int) [][]int {
	avail := [][]int{}
	for i, r := range s {
		for j, v := range r {
			if v == 0 {
				avail = append(avail, []int{i, j})
			}
		}
	}

	return avail
}

func copyBoard(s [][]int) [][]int {
	new := make([][]int, len(s))
	for i := range s {
		row := []int{}
		for j := range s[i] {
			row = append(row, s[i][j])
		}
		new = append(new, row)
	}
	return new
}

func max(s [][]int, maxxer int)

func miniMax(s [][]int, maxPlayer int, maxxing bool) []int {
	maxMove := []int{}
	avail := availSquares(s)

	return maxMove
}
