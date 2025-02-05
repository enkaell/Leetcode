package main

import (
	"fmt"
	"leetcode/array_and_string/mergesortedarray"
	"leetcode/array_and_string/removeduplicatesfromsorted"
	"leetcode/array_and_string/removeelement"
)

func main() {
	mergesortedarray.TestMerge()
	removeelement.TestRemoveElemet()
	removeduplicatesfromsorted.TestRemoveDuplicates()
	removeduplicatesfromsorted.TestRemoveDuplicates2()
	fmt.Println("End")
}
