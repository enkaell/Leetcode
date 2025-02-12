package jumpgame

import (
	"fmt"
)

func canJump(nums []int) bool {
	nums = nums[:len(nums)-1]
	counter := 0
	for _, elem := range nums {
		if elem > counter {
			counter = elem
		}
		counter--
		if (elem == 0) && (counter < 0) {
			return false
		}
	}
	return true
}

func TestCanJump() {
	nums := []int{2, 3, 1, 1, 4}
	if canJump(nums) != true {
		panic("Test case 1 failed")
	}
	nums = []int{3, 2, 1, 0, 4}
	if canJump(nums) != false {
		panic("Test case 2 failed")
	}
	fmt.Println("Test cases passed")
}
