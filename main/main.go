package main

import "github.com/01-edu/z01"

func main () {
	printnbr(-123)
}

func printnbr(n int) {
	var result int
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	for n != 0 {
		result = (result * 10) + (n % 10)
		n /= 10
	}

	for result != 0 {
		z01.PrintRune(rune(result % 10)+'0')
		result /= 10
	}

}

func printcomb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '8'; b++ {
			d := b + 1
			for c := a; c <= '9'; c++ {
				for ; d <= '9'; d++ {

					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(d)

					if a != '9' || b != '8' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
				d = '0'
			}
		}
	} 
}

func printcomb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				z01.PrintRune(k)
				z01.PrintRune(j)
				z01.PrintRune(i)

				if i != '7' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}