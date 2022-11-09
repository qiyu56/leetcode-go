package problems

// 算法不难，细节是魔鬼
// mid +1 or -1, >= or >

// https://leetcode-cn.com/problems/binary-search/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

// LeftBound 1 2 2 2 5
// 找出2的左边界
func LeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right { // 终止条件 left = right + 1
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left >= len(nums) /*都比target小*/ || nums[left] != target /*都比target大或区间内不存在target*/ {
		return -1
	}
	return left
}

func RightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right { // 终止条件 left = right + 1, 循环结束时 nums[left] != target，因为对left的更新一直是mid + 1
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

// 1 2 2 2 5

// SearchRange https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func SearchRange(nums []int, target int) []int {
	left := LeftBound(nums, target)
	right := RightBound(nums, target)
	return []int{left, right}
}
