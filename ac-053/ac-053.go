package ac_053

import "sort"

//查找使用次数
//暴力循环
func search(nums []int, target int) int {
	//m := make(map[int]bool)
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result++
		}
	}
	return result
}

//二分法
func search2(nums []int, target int) int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return 0
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return rightmost - leftmost + 1
}

//总和的差值解法
func missingNumber(nums []int) int {
	l1 := len(nums)
	t1 := 0
	t2 := 0
	for _, v := range nums {
		t1 = t1 + v
	}
	for i := 0; i < l1+1; i++ {
		t2 = t2 + i
	}
	result := t2 - t1
	return result
}

//二分查找
func missingNumber2(nums []int) int {
	left := 0
	right := len(nums)

	for left < right {
		mid := (left + right) >> 1
		if nums[mid] != mid {
			//nums是有序数组，如果mid下标的值和mid的值不相同，说明左边是不连续的数组，如果相同，就在右边找
			right = mid
		} else {
			//缺失的数字就在右边查找，left向上取整 +1
			left = mid + 1
		}
	}
	return left

}