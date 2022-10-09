package golangMinMax

func miniMax(state [][]int, depth int, maxPlayer bool) int {

	avail := availSquares(state)

	terminalCheck, val := isTerminal(state)
	if depth == 0 || terminalCheck {
		return val
	}

	if maxPlayer {
		maxEval := -1
		for _, spot := range avail {
			child := copyBoard(state)
			child[spot[0]][spot[1]] = 1
			eval := miniMax(child, depth-1, false)
			maxEval = max(maxEval, eval)
		}
		return maxEval
	} else {
		minEval := 1
		for _, spot := range avail {
			child := copyBoard(state)
			child[spot[0]][spot[1]] = 2
			eval := miniMax(child, depth-1, true)
			minEval = min(minEval, eval)
		}
		return minEval
	}
}
