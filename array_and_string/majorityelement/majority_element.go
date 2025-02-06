// 169. Majority Element
package majorityelement

import (
	"fmt"
)

func majorityElement1(nums []int) int {
	// Overcomplicated custom solution
	// c := 0
	// half := len(nums)/2
	// if len(nums) <3 {
	//     return nums[0]
	// }
	// for i:=0; i< half; i++ {
	//     for j:=half+i; j < len(nums); j++{
	//         if nums[i] == nums[j] {
	//             nums[i] += nums[i]
	//             c = nums[half+i]
	//             nums[half+i] = nums[j]
	//             nums[j] = c
	//         }
	//     }
	//     }
	// fmt.Println(nums, c)
	// if len(nums) % 2 == 0{ return nums[half+1] } else {return nums[half] }

	// Boyer–Moore majority vote algorithm
	if len(nums) < 3 {
		return nums[0]
	}
	counter := 0
	elem := 0
	for i := 0; i < len(nums); i++ {
		if counter == 0 {
			elem = nums[i]
			counter++
		} else if elem == nums[i] {
			counter++
		} else if elem != nums[i] {
			counter--
		}
	}
	fmt.Println(elem)
	return elem

}

// Using bit manipulation

func majorityElement(nums []int) int {
	n := len(nums)
	majority := 0

	for i := 0; i < 32; i++ { // Iterate over 32-bit integer positions
		bitCount := 0
		mask := 1 << i // Create a mask for the i-th bit
		fmt.Println(mask)
		for _, num := range nums {
			fmt.Println(num & mask)
			if num&mask != 0 { // Check if the i-th bit is set
				bitCount++
			}
		}

		if bitCount > n/2 { // If more than half of the numbers have this bit set
			majority |= mask // Set the bit in the majority element
		}
	}

	return majority
}
