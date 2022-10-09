package golangMinMax

func MiniMax(state [][]int, depth int, maxPlayer bool) int {

	avail := availSquares(state)

	terminalCheck, val := isTerminal(state)
	if depth == 0 || terminalCheck {
		var score int
		if val == 1 {
			score = 1
		} else if val == 2 {
			score = -1
		} else {
			score = 0
		}
		return score
	}

	if maxPlayer {
		maxEval := -1
		for _, spot := range avail {
			child := copyBoard(state)
			child[spot[0]][spot[1]] = 1
			eval := MiniMax(child, depth-1, false)
			maxEval = max(maxEval, eval)
		}
		return maxEval
	} else {
		minEval := 1
		for _, spot := range avail {
			child := copyBoard(state)
			child[spot[0]][spot[1]] = 2
			eval := MiniMax(child, depth-1, true)
			minEval = min(minEval, eval)
		}
		return minEval
	}
}
