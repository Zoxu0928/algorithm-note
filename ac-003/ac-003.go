package ac_003

func findRepeatNumber1(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] {
			return v
		}
		m[v] = true
	}
	return -1
}

// 原地交换， 时间O(n), 空间O(1)
func findRepeatNumber2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// 遍历的索引对应的值等于索引，不动，继续遍历
		for nums[i] != i {
			// 即nums[i] == num[x]
			// 如果i下标对应的值与下标i的元素值做索引的值相等，说明遇到重复值，返回该元素
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			// 不等，交换两值
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	return -1
}
