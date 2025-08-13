package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	mp := make(map[int][]int)
	i := 0
	x := 1
	for i <= len(deck)-1 {

		// fmt.Printf("hello Player %v: %v", x, mp[x])
		if i >= 2 && i%3 == 0 {
			x++
		}
		mp[x] = append(mp[x], deck[i])
		i++
	}
	for k, v := range mp {

		fmt.Printf("Player %v:", k)
		for i, in := range v {
			fmt.Printf(" %v", in)
			if i < len(v)-1 {
				fmt.Printf(",")
			}
		}
		fmt.Printf("\n")

	}
}
