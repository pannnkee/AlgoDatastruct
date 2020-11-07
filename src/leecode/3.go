package leecode

import "fmt"

// 无重复字符的最长子串
func TheLongestSubstringWithoutRepeatedCharacters() {
	str := "abcabcbb"

	maxLen := 0
	strMap := make(map[byte]struct{}, 0)
	for left, right := 0,0; right < len(str); {
		if _, ok := strMap[str[right]]; ok {
			delete(strMap, str[right])
			left++
		} else {
			strMap[str[right]] = struct{}{}
			right++
		}

		if maxLen < right - left {
			maxLen = right - left
		}
	}
	fmt.Println("maxLen:", maxLen)
}































