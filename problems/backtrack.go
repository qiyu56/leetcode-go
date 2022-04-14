package problems

// 回溯问题

// https://leetcode-cn.com/problems/permutations/

var res [][]int

func Permute(nums []int) [][]int {
	res = make([][]int, 0)
	track := make([]int, 0)
	backTrack(nums, track)
	return res
}

func backTrack(nums, track []int) {
	if len(track) == len(nums) {
		backup := make([]int, len(track))
		copy(backup, track)
		res = append(res, backup)
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(track, nums[i]) {
			continue
		}
		track = append(track, nums[i])
		backTrack(nums, track)
		track = track[:len(track)-1]
	}
}

func contains(arr []int, e int) bool {
	for _, val := range arr {
		if val == e {
			return true
		}
	}
	return false
}

var allMethod [][]string

// https://leetcode-cn.com/problems/n-queens/
func SolveNQueens(n int) [][]string {
	allMethod = make([][]string, 0)
	board := make([]int, n) // 第i行皇后放在board[i]列
	for i := 0; i < n; i++ {
		board[i] = -1
	}
	backtrack(board, 0)
	return allMethod
}

func backtrack(board []int, row int) {
	n := len(board)
	if row == n {
		allMethod = append(allMethod, generateBoard(board))
		return
	}
	for col := 0; col < n; col++ {
		if !isValid(board, row, col) {
			continue
		}
		board[row] = col
		backtrack(board, row+1)
		board[row] = -1
	}
}

// (row, col)是否可以放置皇后
// 因为放置时是从上往下逐行放置的，所以只需检查3个方向
func isValid(board []int, row, col int) bool {
	n := len(board)
	// ↑ 方向是否存在冲突
	for i := 0; i < row; i++ {
		if board[i] == col { // 第i行也在第col列放置皇后
			return false
		}
	}
	// ↗ 方向是否存在冲突
	i := row - 1
	j := col + 1
	for i >= 0 && j < n {
		if board[i] == j {
			return false
		}
		i--
		j++
	}
	// ↖ 方向是否存在冲突
	i = row - 1
	j = col - 1
	for i >= 0 && j >= 0 {
		if board[i] == j {
			return false
		}
		i--
		j--
	}
	return true
}

func generateBoard(board []int) []string {
	n := len(board)
	strs := make([]string, n)
	for row := 0; row < n; row++ {
		s := make([]byte, n)
		for col := 0; col < n; col++ {
			if col == board[row] {
				s[col] = 'Q'
			} else {
				s[col] = '.'
			}
		}
		strs[row] = string(s)
	}
	return strs
}
