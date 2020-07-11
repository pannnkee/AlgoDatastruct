package algo

import (
	"fmt"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {

	nums := []int{1,2,3,2}
	fmt.Println(FindRepeatNumber(nums))
}

func TestFindRepeatNumber2(t *testing.T) {
	nums := []int{1,2,3,2}
	fmt.Println(FindRepeatNumber2(nums))
}