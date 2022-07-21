package main

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。
*/

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	//attention  考虑边界条件
	if target > nums[len(nums)-1] {
		return len(nums)
	} else if target == nums[len(nums)-1] {
		return len(nums) - 1
	}
	if target < nums[0] {
		return 0
	} else if target == nums[0] {
		return 0
	}

	//二分查找

	for end > start+1 {
		mid := 0
		if end+start%2 == 0 {
			mid = (end + start) / 2
		} else {
			mid = (end + start + 1) / 2
		}
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else {
			return mid
		}
	}
	return end
}
