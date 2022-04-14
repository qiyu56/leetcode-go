package problems

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
