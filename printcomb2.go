package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for k := '0'; k <= '9'; k++ {
			j := k + 1

			for a := i; a <= '9'; a++ {
				for ; j <= '9'; j++ {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(k))
					z01.PrintRune(' ')
					z01.PrintRune(rune(a))
					z01.PrintRune(rune(j))

					if i != '9' || k != '8' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				j = '0'
			}
		}
	}

	z01.PrintRune('\n')
	return
}
