package problems

// https://leetcode-cn.com/problems/rotate-image/
func rotate(matrix [][]int) {
	n := len(matrix)
	// 先沿(0,0) -> (n,n)旋转
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 再水平翻转
	for _, arr := range matrix {
		reverseRow(arr)
	}
}

func reverseRow(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

// https://leetcode-cn.com/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {
	h, w := len(matrix), len(matrix[0])
	top, left, right, bottom := 0, 0, w-1, h-1
	res := make([]int, 0)
	for len(res) < h*w {
		// left -> right
		if top <= bottom {
			for i := left; i <= right; i++ {
				res = append(res, matrix[top][i])
			}
			top++
		}
		// top -> bottom
		if right >= left {
			for i := top; i <= bottom; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}
		// right -> left
		if bottom >= top {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}
		// bottom -> top
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}

// https://leetcode-cn.com/problems/spiral-matrix-ii/
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	top, left, right, bottom := 0, 0, n-1, n-1
	num := 1
	for num <= n*n {
		if top <= bottom {
			for i := left; i <= right; i++ {
				res[top][i] = num
				num++
			}
			top++
		}
		if right >= left {
			for i := top; i <= bottom; i++ {
				res[i][right] = num
				num++
			}
			right--
		}
		if bottom >= top {
			for i := right; i >= left; i-- {
				res[bottom][i] = num
				num++
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				res[i][left] = num
				num++
			}
			left++
		}
	}
	return res
}
