package algo

import "sort"

//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
//数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
func FindRepeatNumber(nums []int) int {
	var flag int
	for k, v := range nums {
		flag = v
		for _, value := range nums[k+1:] {
			if flag == value {
				return flag
			}
		}
	}
	return -1
}

func FindRepeatNumber2(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

func FindRepeatNumber3(nums []int) int {
	m := map[int]int{}

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v] = 1
		}
	}
	return -1
}
