package removeduplicatesfromsorted

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var k = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}

	}
	return k + 1
}

func TestRemoveDuplicates() {
	k := removeDuplicates([]int{1, 1, 2})
	if k != 2 {
		panic("Test case 1 failed")
	}
	k = removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	if k != 5 {
		panic("Test case 2 failed")
	}
	fmt.Println("Test cases passed")
}

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	var k = 1
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}

func TestRemoveDuplicates2() {
	k := removeDuplicates2([]int{1, 1, 1, 2, 2, 3})
	if k != 5 {
		panic("Test case 1 failed")
	}
	k = removeDuplicates2([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
	if k != 7 {
		panic("Test case 2 failed")
	}
	fmt.Println("Test cases passed")
}
