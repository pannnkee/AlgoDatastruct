package leecode

import "sort"

// 给定一个无序的整数数组，找到其中最长上升子序列的长度。
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	var temp int
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		//找前面所有比他小的数的dp
		temp = 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if temp < dp[j] {
					temp = dp[j]
				}
			}
		}
		dp[i] = temp + 1
	}
	sort.Ints(dp)
	return dp[len(dp) - 1]
}
