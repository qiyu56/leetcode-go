package problems

import "LeetCode/common"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := common.Queue{}
	q.Push(root)
	depth := 1
	for q.Len() > 0 {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			cur := q.Pop().(*TreeNode)
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				q.Push(cur.Left)
			}
			if cur.Right != nil {
				q.Push(cur.Right)
			}
		}
		depth++
	}
	return depth
}

// https://leetcode-cn.com/problems/open-the-lock/
func OpenLock(deadends []string, target string) int {
	deads := make(map[string]bool)
	for _, s := range deadends {
		deads[s] = true
	}
	visited := make(map[string]bool)
	q := common.Queue{}
	q.Push("0000")
	visited["0000"] = true
	step := 0
	for q.Len() > 0 {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			cur := q.Pop().(string)
			if _, ok := deads[cur]; ok {
				continue
			}
			if cur == target {
				return step
			}

			for j := 0; j < 4; j++ {
				up, down := plusOne(cur, j), minusOne(cur, j)
				if _, ok := visited[up]; !ok {
					q.Push(up)
					visited[up] = true
				}
				if _, ok := visited[down]; !ok {
					q.Push(down)
					visited[down] = true
				}
			}
		}
		step++
	}
	return -1
}

func plusOne(s string, i int) string {
	bytes := []byte(s)
	if bytes[i] == '9' {
		bytes[i] = '0'
	} else {
		bytes[i] += 1
	}
	return string(bytes)
}

func minusOne(s string, i int) string {
	bytes := []byte(s)
	if bytes[i] == '0' {
		bytes[i] = '9'
	} else {
		bytes[i] -= 1
	}
	return string(bytes)
}
