package main

import (
	"os"

	"github.com/01-edu/z01"
)

// var (
// 	minus1 bool
// 	minus2 bool
// )

// func main() {
// 	arg := os.Args[1:]

// 	int1 := 0
// 	int2 := 0

// 	if len(arg) == 3 {
// 		if !(IsNumm(arg[0]) && IsNumm(arg[2])) {
// 			return
// 		}

// 		number1 := arg[0]
// 		number2 := arg[2]

// 		for i, v := range number1 {
// 			if i == 0 && v == '-' {
// 				minus1 = true
// 				continue
// 			}
// 			intNumber1 := int(v - '0')
// 			int1 = int1*10 + intNumber1
// 			if minus1 {
// 				int1 *= -1
// 			}
// 		}

// 		for i, v := range number2 {
// 			if i == 0 && v == '-' {
// 				minus2 = true
// 				continue
// 			}
// 			intNumber2 := int(v - '0')
// 			int2 = int2*10 + intNumber2
// 			if minus2 {
// 				int2 *= -1
// 			}
// 		}

// 		if int1 < -9223372036854775800 || int1 > 9223372036854775800 {
// 			return
// 		}
// 		if int2 < -9223372036854775800 || int2 > 9223372036854775800 {
// 			return
// 		}

// 		if arg[1] == "+" {

// 			result := int1 + int2

// 			Itoa(result)
// 			z01.PrintRune('\n')
// 		}
// 		if arg[1] == "-" {
// 			result := int1 - int2
// 			if result < 0 {
// 				result *= -1
// 				z01.PrintRune('-')
// 			}
// 			Itoa(result)
// 			z01.PrintRune('\n')
// 		}
// 		if arg[1] == "*" {
// 			result := int1 * int2

// 			if (minus2 || minus1) && !(minus1 && minus2) {
// 				z01.PrintRune('-')
// 			}

// 			Itoa(result)
// 			z01.PrintRune('\n')
// 		}
// 		if arg[1] == "/" {
// 			if int2 == 0 {
// 				Message := "No division by 0"
// 				for _, v := range Message {
// 					z01.PrintRune(rune(v))
// 				}
// 				z01.PrintRune('\n')
// 				return
// 			}
// 			result := int1 / int2

// 			if (minus2 || minus1) && !(minus1 && minus2) {
// 				z01.PrintRune('-')
// 			}

// 			Itoa(result)
// 			z01.PrintRune('\n')
// 		}
// 		if arg[1] == "%" {
// 			if int2 == 0 {
// 				Message := "No modulo by 0"
// 				for _, v := range Message {
// 					z01.PrintRune(rune(v))
// 				}
// 				z01.PrintRune('\n')
// 				return
// 			}
// 			result := int1 % int2

// 			if (minus2 || minus1) && !(minus1 && minus2) {
// 				z01.PrintRune('-')
// 			}

// 			Itoa(result)
// 			z01.PrintRune('\n')
// 		}

// 	} else {
// 		return
// 	}
// }

// func Itoa(x int) {
// 	if x == 0 {
// 		return
// 	}
// 	Itoa(x / 10)
// 	z01.PrintRune(rune(x%10 + 48))
// }

// func IsNumm(s string) bool {
// 	for _, v := range s {
// 		if !(v >= 48 && v <= 57 || v == 45) {
// 			return false
// 		}
// 	}
// 	return true
// }
func Plus(a, b int) int {
	return a + b
}

func Mult(a, b int) int {
	return a * b
}

func Minus(a, b int) int {
	return a - b
}

func Divid(a, b int) int {
	return a / b
}

func Modulo(a, b int) int {
	return a % b
}

func JustDoIt(f func(int, int) int, a, b int) {
	PrintInt(f(a, b))
	z01.PrintRune('\n')
}

func solo(n, dot int) {
	if n == 0 {
		return
	}
	digit := int((n % 10) * dot)
	n /= 10
	solo(n, dot)
	z01.PrintRune(rune(digit) + 48)
}

func PrintInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	dot := 1
	if n < 0 {
		z01.PrintRune('-')
		dot *= -1
	}
	solo(n, dot)
}

func PrintMessg(x string) {
	for _, ch := range x {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func ABS(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func AtoiX(s string) int {
	result := 0
	count := 0
	if s == "" {
		return 0
	}
	for _, letter := range s {
		count++
		letter++
	}
	flag := 0
	plus := 0
	if s[0] == '-' {
		flag = 1
	} else if s[0] == '+' {
		flag = 1
		plus = 1
	}
	for _, ch := range s[flag:count] {
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + int(ch-'0')
	}
	if flag == 1 && plus == 0 {
		return -result
	}
	return result
}

func IsNumm(s string) bool {
	for _, ch := range s {
		if !(ch >= 48 && ch <= 57 || ch == 45) {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	if !(IsNumm(args[0]) && IsNumm(args[2])) {
		return
	}
	a, sign, b := AtoiX(args[0]), args[1], AtoiX(args[2])
	if ABS(a) > 922337203685477580 || ABS(b) > 922337203685477580 {
		return
	}
	switch sign {
	case "+":
		JustDoIt(Plus, a, b)
	case "-":
		JustDoIt(Minus, a, b)
	case "*":
		JustDoIt(Mult, a, b)
	case "/":
		if b == 0 {
			PrintMessg("No division by 0")
			return
		}
		JustDoIt(Divid, a, b)
	case "%":
		if b == 0 {
			PrintMessg("No modulo by 0")
			return
		}
		JustDoIt(Modulo, a, b)
	default:
		return
	}
}
