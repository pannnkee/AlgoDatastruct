package leecode

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转
func Reverse(x int) int {
	var res int
	for x != 0 {
		if temp := int32(res); temp != (temp*10)/10 {
			return 0
		}
		res = res *10 + x %10
		x = x / 10
	}
	return res
}

























