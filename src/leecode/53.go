package leecode

import (
	"sort"
)

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
func MaxSubArray(nums []int) int {
	dp := make([]int, len(nums))

	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = nums[i] + dp[i-1]
		}
	}
	sort.Ints(dp)
	return dp[len(nums)-1]
}

func A(nums []int) int {
	temp := -2147483648
	sum := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum = 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			if sum > temp {
				temp = sum
			}
		}
	}
	return  temp
}
