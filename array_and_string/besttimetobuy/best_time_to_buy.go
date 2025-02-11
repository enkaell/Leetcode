package besttimetobuy

import (
	"fmt"
)

func maxProfit(prices []int) int {
	profit := 0
	min := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}
	fmt.Println(profit)
	//     Didnt work
	//     if prices[i] < min && i < maxIndex {
	//         min = prices[i]
	//         minIndex = i
	//         fmt.Println("Min is set", min)
	//         }
	//     if prices[len(prices)-i-1] > max &&  len(prices)-i-1 > minIndex{
	//         max = prices[len(prices)-i-1]
	//         maxIndex = len(prices)-i-1
	//         fmt.Println("Max is set", max, maxIndex, minIndex)
	//         }
	//     if (max-min>profit){
	//         fmt.Println(i, min,max)
	//         profit = max-min
	//     }
	// }
	// Worst use-case scenario: nested loops O(n^2) complexity
	// for i:=0; i<len(prices)-1; i++{
	//     if prices[i] < prices[i+1]{
	//     for j:=i+1; j<len(prices); j++ {
	//         if prices[j] - prices[i] > profit {
	//             profit = prices[j]-prices[i]
	//         }
	//     }
	//     }
	// }
	return profit
}

func TestMaxProfit() {
	prices := []int{7, 1, 5, 3, 6, 4}
	if maxProfit(prices) != 5 {
		panic("Test case 1 failed")
	}
	prices = []int{7, 6, 4, 3, 1}
	if maxProfit(prices) != 0 {
		panic("Test case 2 failed")
	}
	prices = []int{2, 4, 1}
	if maxProfit(prices) != 2 {
		panic("Test case 3 failed")
	}
}

func maxProfit2(prices []int) int {
	profit1 := 0
	elem := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] > elem {
			profit1 += prices[i] - elem
		}
		elem = prices[i]
	}
	return profit1
}

func TestMaxProfit2() {
	prices := []int{7, 1, 5, 3, 6, 4}
	if maxProfit2(prices) != 7 {
		panic("Test case 1 failed")
	}
	prices = []int{7, 6, 4, 3, 1}
	if maxProfit2(prices) != 0 {
		panic("Test case 2 failed")
	}
	prices = []int{2, 4, 1}
	if maxProfit2(prices) != 2 {
		panic("Test case 3 failed")
	}
}
