package problems

// MinWindow https://leetcode-cn.com/problems/minimum-window-substring/
func MinWindow(s string, t string) string {
	need, window := make(map[rune]int), make(map[rune]int)
	// need: t中字符的个数
	// window: 窗口中t中字符的个数
	runeS, runeT := []rune(s), []rune(t)
	for _, c := range runeT {
		need[c] += 1
	}
	left, right, valid := 0, 0, 0 // [left, right) , valid: 子串中满足t中字符的个数
	start, length := 0, len(s)+1  // 最小子串的start, len
	for right < len(runeS) {
		c := runeS[right]
		// 窗口扩大
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] { // 窗口中c的个数等于t中c的个数
				valid++
			}
		}
		// window中包含了t中的所有字符，尝试从左侧缩小窗口
		for valid == len(need) {
			// 更新最小覆盖子串
			if right-left < length {
				start = left
				length = right - left
			}
			c := runeS[left]
			left++
			if _, ok := need[c]; ok {
				if window[c] == need[c] { // 边界条件，如果再删掉一个c后就不满足条件了
					valid--
				}
				window[c]--
			}
		}
	}
	if length == len(s)+1 {
		return ""
	} else {
		return s[start : start+length]
	}
}

// https://leetcode-cn.com/problems/permutation-in-string/submissions/
func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[rune]int), make(map[rune]int)
	rune1, rune2 := []rune(s1), []rune(s2)
	for _, c := range rune1 {
		need[c]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(rune2) {
		c := rune2[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(rune1) { // s1的排列之一是s2的子串，子串肯定是连续的
			if valid == len(need) {
				return true
			}
			c := rune2[left]
			left++
			if _, ok := need[c]; ok {
				if window[c] == need[c] {
					valid--
				}
				window[c]--
			}
		}
	}
	return false
}

// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	for _, c := range []byte(p) {
		need[c]++
	}
	left, right := 0, 0
	valid := 0
	res := make([]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			c := s[left]
			left++
			if _, ok := need[c]; ok {
				if window[c] == need[c] {
					valid--
				}
				window[c]--
			}
		}
	}
	return res
}

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for window[c] > 1 {
			t := s[left]
			left++
			window[t]--
		}
		tmp := right - left
		if tmp > res {
			res = tmp
		}
	}
	return res
}

func minWindow1(s, p string) string {
	needs := make(map[byte]int)
	window := make(map[byte]int)
	for _, c := range p {
		needs[byte(c)]++
	}
	left, right := 0, 0 // [left, right)
	valid := 0
	length := len(s) + 1
	start := 0
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := needs[c]; ok {
			window[c]++
			if window[c] == needs[c] {
				valid++
			}
		}
		// 先向右扩展，直至满足条件 right++
		// 再从左侧开始缩小，寻找最优结果 left++
		for len(needs) == valid {
			if right-left < length {
				length = right - left
				start = left
			}
			d := s[left]
			left++
			if _, ok := needs[d]; ok {
				if window[d] == needs[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if length == len(s)+1 {
		return ""
	}
	return s[start : start+length]
}

func lengthOfLongestSubstring1(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0 // [left, right)
	result := 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		if window[c] > 1 { // 窗口中出现了重复元素，收缩
			d := s[left]
			left++
			window[d]--
		}
		if right-left > result {
			result = right - left
		}
	}
	return result
}
