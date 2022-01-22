package ac_003

import (
	"fmt"
	"testing"
)

func TestAC003(t *testing.T) {
	nums := []int{2, 3, 1, 2, 9, 0, 4, 5}
	number2 := findRepeatNumber2(nums)
	fmt.Println(number2)
}
