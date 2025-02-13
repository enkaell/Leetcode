// 55. Jump Game
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

// 45. Jump Game II
func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	if nums[0] >= len(nums)-1 {
		return 1
	}
	// Unoptimized O(n^2) solution
	count := 0
	upperBound := 0
	lowerBound := 0
	for upperBound < len(nums)-1 {
		maxElemInWindow := 0
		fmt.Println("We are iterating over the batch: ", nums[lowerBound:upperBound+1])
		for i := lowerBound; i < upperBound+1; i++ {
			if nums[i]+i > maxElemInWindow {
				maxElemInWindow = nums[i] + i
			}
		}
		lowerBound = upperBound
		upperBound = maxElemInWindow
		count++
	}

	return count
}

func TestJump() {
	nums := []int{2, 3, 1, 1, 4}
	if jump(nums) != 2 {
		panic("Test case 1 failed")
	}
	nums = []int{2, 3, 0, 1, 4}
	if jump(nums) != 2 {
		panic("Test case 2 failed")
	}
	fmt.Println("Test cases passed")
}
