// 88. Merge Sorted Array
package mergesortedarray

import (
	"fmt"
	"leetcode/array_and_string/utils"
	"sort"
)

func Merge(nums1 []int, m int, nums2 []int, n int) {
	var i = m - 1
	var j = n - 1
	var index = m + n - 1
	for j >= 0 {
		if i >= 0 && (nums1[i] > nums2[j]) {
			nums1[index] = nums1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
		index--
	}
}
func TestMerge() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	testCase := nums1
	sort.Ints(testCase)
	Merge(nums1, 3, []int{2, 5, 6}, 3)
	if !utils.CompareSlices(nums1, testCase) {
		fmt.Println("Test case 1 failed")
	}

	nums1 = []int{1}
	testCase = nums1
	sort.Ints(testCase)
	Merge(nums1, 1, []int{}, 0)
	if !utils.CompareSlices(nums1, testCase) {
		fmt.Println("Test case 2 failed")
	}

	nums1 = []int{0}
	testCase = nums1
	sort.Ints(testCase)
	Merge(nums1, 0, []int{1}, 1)
	if !utils.CompareSlices(nums1, testCase) {
		fmt.Println("Test case 3 failed")
	}
	fmt.Println("Test cases passed")

}
