package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for a := '9'; a >= '0'; a-- {
		for b := '9'; b >= '0'; b-- {

			d := b - 1

			for c := a; c >= '0'; c-- {
				for ; d >= '0'; d-- {

					z01.PrintRune(a)
					z01.PrintRune(b)

					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(d)

					if b != '1' || a != '0' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				d = '9'
			}
		}
	}
}
