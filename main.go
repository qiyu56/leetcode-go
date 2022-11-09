package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums2))
	}
	res := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if i == 0 || j == 0 {
				if nums1[i] == nums2[j] {
					dp[i][j] = 1
				}
			} else {
				if nums1[i] == nums2[j] {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		tmp := IntMin(height[left], height[right]) * (right - left)
		res = IntMax(res, tmp)

		if height[left] < height[right] {
			curLeft := height[left]
			for left < right && height[left] <= curLeft {
				left++
			}
		} else {
			curRight := height[right]
			for right > left && height[right] <= curRight {
				right--
			}
		}
	}
	return res
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for ; index < length && str1[index] == str2[index]; index++ {
	}
	return str1[:index]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	window := make(map[byte]int)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := s[right]
		right += 1
		window[c]++
		for window[c] > 1 {
			t := s[left]
			left++
			window[t]--
		}
		res = max(right-left, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func validUtf8(data []int) bool {
	n := 1
	byteCount := 0
	for _, num := range data {
		s := fmt.Sprintf("%08b", num)
		if s[0] == '0' && n == 1 {
			continue
		}
		if s[:2] == "11" {
			if byteCount > 0 {
				return false
			}
			i := 0
			for s[i] == '1' {
				i += 1
			}
			if i > 4 {
				return false
			}
			n = i
			byteCount += 1
		} else if s[:2] == "10" && byteCount < n && byteCount > 0 {
			byteCount += 1
		} else {
			return false
		}
		if byteCount == n {
			n = 1
			byteCount = 0
		}
	}
	return n == 1 && byteCount == 0
}

func nextGreaterElement(nums []int) []int {
	ret := make([]int, len(nums))
	s := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		for len(s) > 0 && s[len(s)-1] <= nums[i] {
			s = s[:len(s)-1]
		}
		if len(s) > 0 {
			ret[i] = s[len(s)-1]
		} else {
			ret[i] = -1
		}
		s = append(s, nums[i])
	}
	return ret
}

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	ret := 1
	cur := pairs[0][1]
	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] > cur {
			ret += 1
			cur = pairs[i][1]
		}
	}
	return ret
}

func finalPrices(prices []int) []int {
	ret := make([]int, len(prices))
	s := make([]int, 0)
	for i := len(prices) - 1; i >= 0; i-- {
		for len(s) > 0 && prices[i] < s[len(s)-1] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			ret[i] = prices[i]
		} else {
			ret[i] = prices[i] - s[len(s)-1]
		}
		s = append(s, prices[i])
	}
	return ret
}

func canJump(nums []int) bool {
	var f func(idx int) bool
	f = func(idx int) bool {
		for i := idx - 1; i >= 0; i-- {
			step := idx - i
			if nums[i] >= step {
				if i == 0 {
					return true
				}
				return f(i)
			}
		}
		return false
	}
	return f(len(nums) - 1)
}

func canJump1(nums []int) bool {
	k := 0
	for i := 0; i < len(nums); i++ {
		if k < i {
			return false
		}
		if i+nums[i] > k {
			k = i + nums[i]
		}
	}
	return true
}

func jump(nums []int) int {
	start, end := 0, nums[0]
	ans := 0
	for end < len(nums) {
		maxPos := end
		for i := start; i < end; i++ {
			if i+nums[i] > maxPos {
				maxPos = i + nums[i]
			}
		}
		start = end
		end = maxPos
		ans += 1
	}
	return ans
}

func fizzBuzz(n int) []string {
	var ans []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans = append(ans, "FizzBuzz")
		} else if i%3 == 0 {
			ans = append(ans, "Fizz")
		} else if i%5 == 0 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, strconv.Itoa(i))
		}
	}
	return ans
}

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var record []int
	sum := 0
	var dfs func(idx int)
	dfs = func(num int) {
		if sum == n && len(record) == k {
			dst := make([]int, k)
			copy(dst, record)
			ans = append(ans, dst)
			return
		}
		if sum > n || len(record) >= k {
			return
		}

		for i := num; i <= 9; i++ {
			record = append(record, i)
			sum += i
			dfs(i + 1)
			record = record[:len(record)-1]
			sum -= i
		}
	}
	dfs(1)
	return ans
}

func numSpecial(mat [][]int) int {
	r := make([]int, len(mat))
	c := make([]int, len(mat[0]))

	for i, row := range mat {
		for j, _ := range row {
			r[i] += mat[i][j]
			c[j] += mat[i][j]
		}
	}

	ret := 0
	for i, row := range mat {
		for j, _ := range row {
			if mat[i][j] == 1 && r[i] == 1 && c[j] == 1 {
				ret += 1
			}
		}
	}
	return ret
}

func reorderSpaces(text string) string {
	var words []string
	w := ""
	whiteSpaceCount := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			whiteSpaceCount++
			if len(w) > 0 {
				words = append(words, w)
				w = ""
			}
		} else {
			w = w + string(text[i])
		}
	}
	if len(w) > 0 {
		words = append(words, w)
	}
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", whiteSpaceCount)
	}

	whiteSpace := strings.Repeat(" ", whiteSpaceCount/(len(words)-1))
	ans := strings.Join(words, whiteSpace)
	ans += strings.Repeat(" ", whiteSpaceCount%(len(words)-1))
	return ans
}

func main() {
	fmt.Println(reorderSpaces("  this   is  a sentence "))

}
