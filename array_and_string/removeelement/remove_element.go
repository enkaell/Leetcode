// 27. Remove Element
package removeelement

import "fmt"

func removeElement(nums []int, val int) int {
	var k = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func TestRemoveElemet() {
	nums := []int{3, 2, 2, 3}
	val := 3
	k := removeElement(nums, val)
	if k != 2 {
		panic("Test case 1 failed")
	}

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	k = removeElement(nums, val)
	if k != 5 {
		panic("Test case 2 failed")
	}
	fmt.Println("Test cases passed")
}
