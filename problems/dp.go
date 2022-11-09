package problems

import "sort"

// https://leetcode-cn.com/problems/coin-change/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // 当目标金额为i时，至少需要dp[i]枚硬币凑出
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	// base case
	dp[0] = 0
	// 依次计算所有子问题，得出最终解
	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// LengthOfLIS https://leetcode-cn.com/problems/longest-increasing-subsequence/
func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j] > max {
				max = dp[j]
			}
		}
		dp[i] = max + 1
	}
	res := 1
	for i := 0; i < len(dp); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// MaxEnvelopes https://leetcode-cn.com/problems/russian-doll-envelopes/
func MaxEnvelopes(envelopes [][]int) int {
	// w相同时，按h降序排列，则这些降序排列的h中最多只会有一个被选入递增子序列
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return false
	})

	// 在h数组中寻找最长递增子序列
	//dp := make([]int, len(envelopes))
	//dp[0] = 1
	//res := 1
	//for i := 1; i < len(envelopes); i++ {
	//	max := 0
	//	for j := 0; j < i; j++ {
	//		if envelopes[j][1] < envelopes[i][1] && dp[j] > max {
	//			max = dp[j]
	//		}
	//	}
	//	dp[i] = max + 1
	//	if dp[i] > res {
	//		res = dp[i]
	//	}
	//}
	//return res

	top := make([]int, len(envelopes))
	piles := 0

	for i := 0; i < len(top); i++ {
		poker := envelopes[i][1]
		left, right := 0, piles
		for left < right {
			mid := left + (right-left)/2
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == piles {
			piles++
		}
		top[left] = poker
	}
	return piles
}

func lengthOfLITBS(nums []int) int {
	top := make([]int, len(nums))
	piles := 0

	for i := 0; i < len(nums); i++ {
		poker := nums[i]
		left, right := 0, piles
		for left < right {
			mid := left + (right-left)/2
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == piles {
			piles++
		}
		top[left] = poker
	}
	return piles
}

// 0-1背包
func knapsack(v int, n int, vw [][]int) int {
	dp := make([][]int, v+1)
	for i := 0; i <= v; i++ {
		dp[i] = make([]int, n+1)
	}
	// dp[i][j] 背包最大体积为i, 有前j个物品时最多可以装多大重量
	for i := 0; i <= v; i++ {
		for j := 1; j <= n; j++ {
			cv := vw[j-1][0]
			cw := vw[j-1][1]
			if i-cv >= 0 { // 有体积装，且没有装第j件物品时
				dp[i][j] = max(dp[i-cv][j-1]+cw, dp[i][j-1])
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[v][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func BackPack(m int, a []int) int {
	dp := make([]int, m+1)
	for i := 0; i < len(a); i++ {
		for j := m; j >= a[i]; j-- {
			dp[j] = max(dp[j], dp[j-a[i]]+a[i])
		}
	}
	return dp[m]
}

func BackPackII(m int, a []int, v []int) int {
	dp := make([]int, m+1)
	for i := 0; i < len(a); i++ {
		for j := m; j >= a[i]; j-- {
			dp[j] = max(dp[j], dp[j-a[i]]+v[i])
		}
	}
	return dp[m]
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}
	res := 0
	for _, val := range dp {
		res = max(res, val)
	}
	return res
}

func maxEnvelopes(envelopes [][]int) int {
	/**
	  因为两个w相同的信封不能相互包含，w相同时将h逆序排序，则这些逆序h中最多只会有一个被选入递增子序列，
	  保证了最终的信封序列中不会出现w相同的情况。
	*/
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] > envelopes[j][0] {
			return false
		} else {
			return envelopes[i][1] > envelopes[j][1]
		}
	})
	nums := make([]int, len(envelopes))
	for i := 0; i < len(nums); i++ {
		nums[i] = envelopes[i][1]
	}
	return lengthOfLIS(nums)
}
