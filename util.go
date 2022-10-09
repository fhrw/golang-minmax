package golangMinMax

func isTerminal(state [][]int) (bool, int) {

	// check rows
	for _, r := range state {
		isFull := full(r)
		sameCheck, val := allSame(r)
		if isFull && sameCheck {
			return true, val
		}
	}

	// check cols
	for i := range state {
		col := getCol(state, i)
		isFull := full(col)
		sameCheck, val := allSame(col)
		if isFull && sameCheck {
			return true, val
		}
	}

	//check diags
	diags := [][]int{getDiag(state, true), getDiag(state, false)}
	for _, d := range diags {
		isFull := full(d)
		sameCheck, val := allSame(d)
		if isFull && sameCheck {
			return true, val
		}
	}

	// check for draw
	if len(availSquares(state)) == 0 {
		return true, 0
	}

	return false, 0

}

func getDiag(state [][]int, direction bool) []int {

	var idx int
	s := []int{}

	if direction {
		idx = 0
		for _, r := range state {
			s = append(s, r[idx])
			idx++
		}
	} else {
		idx = len(state)
		for _, r := range state {
			s = append(s, r[idx])
			idx--
		}
	}

	return s
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

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
