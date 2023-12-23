package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	var card int = 0

	for i := 1; i <= 4; i++ {
		fmt.Printf("Player %d:", i)

		for ; card < 12; card++ {
			if card != 2 && card != 5 && card != 8 && card != 11 {
				fmt.Printf(" %d,", deck[card])
			}

			if card == 2 || card == 5 || card == 8 || card == 11 {
				fmt.Printf(" %d", deck[card])
				break
			}
		}
		card++
		z01.PrintRune('\n')
	}
}
