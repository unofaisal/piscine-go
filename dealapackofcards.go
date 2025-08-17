package piscine

import "fmt"

// 12 cards / 4 pple

func DealAPackOfCards(deck []int) {
	for i := 1; i <= 3; {
		fmt.Printf("Player %v: ", i)
		for x := 1; x <= 12; x++ {
			fmt.Printf("%v", x)

			if x != 12 {
				if x%3 == 0 {
					i++
					fmt.Printf("\n")
					fmt.Printf("Player %v: ", i)
				} else {
					fmt.Printf(", ")
				}
			}
		}
	}
	fmt.Printf("\n")
}
