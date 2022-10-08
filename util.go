package golangMinMax

func isTerminal(state [][]int) (bool, int) {

	for _, r := range state {
		full := full(r)
		sameCheck, val := allSame(r)
		if full && sameCheck {
			return true, val
		}
	}

	return false, -1
}

func allSame(s []int) (bool, int) {
	first := s[0]
	rest := s[1:]

	for i := range rest {
		if rest[i] != first {
			return false, -1
		}
	}

	return true, first
}

func full(s []int) bool {

	for _, v := range s {
		if v == 0 {
			return false
		}
	}

	return true
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
