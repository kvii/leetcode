package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	var M, N int
	if M = len(matrix); M == 0 {
		return false
	}
	if N = len(matrix[0]); N == 0 {
		return false
	}

	m, n := 0, N-1
	for m < M && n >= 0 {
		switch i := matrix[m][n]; {
		case target > i:
			m++
		case target < i:
			n--
		case target == i:
			return true
		}
	}
	return false
}
