// 189. Rotate Array
package rotatearray

func rotate(nums []int, k int) {
	if k != 0 {
		// c := 0
		// count := 0
		// arr := []int{}
		// nums = append(nums[len(nums)-k:],nums[:len(nums)-k]...)
		// for i:=len(nums)-k-1; i>=0 ; i-- {
		//     c = nums[i]
		//     nums[i] = nums[i+k]
		//     nums[i+k] = c
		//     count++
		//     fmt.Println(nums)
		// }
		// k = k % (len(nums))
		// for i:=0; i<k; i++ {
		//     arr = append(arr, nums[len(nums)-1])
		// }
		// for i:=len(nums)-k; i<len(nums); i++{
		//     arr = append(arr,nums[i])
		// }
		// for i:=0; i<len(nums)-k;i++{
		//     arr = append(arr,nums[i])
		// }
		// nums = append(arr, nums[:len(nums)-1]...)
		// fmt.Println(arr)
		// copy(nums,arr)
		// fmt.Println(nums)
		k = k % len(nums)
		copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
	}

}
