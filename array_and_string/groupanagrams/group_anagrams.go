package groupanagrams

import (
	"fmt"
	"sort"
)

func formatStringToSortedKey(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[string]int)
	res := [][]string{}

	for _, elem := range strs {
		elemASCIIArray := formatStringToSortedKey(elem)

		index, ok := hashMap[elemASCIIArray]
		if ok {
			res[index] = append(res[index], elem)
		} else {
			elIndex := len(hashMap)
			hashMap[elemASCIIArray] = elIndex
			res = append(res, []string{elem})
			fmt.Println(elemASCIIArray, elem)
		}
	}
	return res
}
