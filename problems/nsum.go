package problems

import "sort"

// TwoSum nums中有且只有两个数和为target，找出这两个数
func TwoSum(nums []int, target int) []int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	// 双指针
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{nums[left], nums[right]}
		} else if sum > target {
			right--
		} else if sum < target {
			left++
		}
	}
	return []int{}
}

// TwoSumTarget 存在多对和为target的元素，返回所有元素对，不能重复
func TwoSumTarget(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	return twoSumTarget(nums, 0, target)
}

func twoSumTarget(nums []int, start int, target int) [][]int {
	// 双指针
	left, right := start, len(nums)-1
	res := make([][]int, 0)
	for left < right {
		l, r := nums[left], nums[right]
		sum := l + r
		if sum == target {
			res = append(res, []int{l, r})
			// 排除重复的
			for left < right && nums[left] == l {
				left++
			}
			for left < right && nums[right] == r {
				right--
			}
		} else if sum > target {
			for left < right && nums[right] == r {
				right--
			}
		} else if sum < target {
			for left < right && nums[left] == l {
				left++
			}
		}
	}
	return res
}

// ThreeSum https://leetcode-cn.com/problems/3sum/
func ThreeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		tuples := twoSumTarget(nums, i+1, 0-nums[i])
		for _, tuple := range tuples {
			tuple = append(tuple, nums[i])
			res = append(res, tuple)
		}
		// 排除重复，i落在重复元素的最后一个，因为外层for循环还有i++
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

// NSum 调用前先给nums排序，因为内部存在递归调用
func NSum(nums []int, n int, start int, target int) [][]int {
	res := make([][]int, 0)
	if n < 2 || n > len(nums) {
		return res
	}
	if n == 2 {
		left, right := start, len(nums)-1
		for left < right {
			l, r := nums[left], nums[right]
			sum := l + r
			if sum == target {
				res = append(res, []int{l, r})
				// 排除重复的
				for left < right && nums[left] == l {
					left++
				}
				for left < right && nums[right] == r {
					right--
				}
			} else if sum > target {
				for left < right && nums[right] == r {
					right--
				}
			} else if sum < target {
				for left < right && nums[left] == l {
					left++
				}
			}
		}
	} else {
		for i := start; i <= len(nums)-n; i++ {
			sub := NSum(nums, n-1, i+1, target-nums[i])
			for _, arr := range sub {
				arr = append(arr, nums[i])
				res = append(res, arr)
			}
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}
