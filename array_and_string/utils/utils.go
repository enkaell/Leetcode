package utils

func CompareSlices(nums1 []int, testCase []int) bool {
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != testCase[i] {
			return false
		}
	}
	return true
}
