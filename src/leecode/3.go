package leecode

import "fmt"

// 无重复字符的最长子串
func TheLongestSubstringWithoutRepeatedCharacters() {
	s := "pwwkew"

	maxLen := 0
	mapStr := make(map[byte]struct{}, 0)
	for left, right := 0,0; right<len(s); {
		if _, ok := mapStr[s[right]]; ok {
			delete(mapStr, s[left])
			left++
		} else {
			mapStr[s[right]] = struct{}{}
			right++
		}
		if maxLen < right - left {
			maxLen = right - left
		}
	}
	fmt.Println("maxLen:", maxLen)
}































