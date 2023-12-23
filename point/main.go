package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := point{}

	setPoint(&points)

	slice := []int{}
	slice2 := []int{}

	for points.x != 0 {
		slice = append(slice, points.x%10)
		points.x /= 10
	}

	for points.y != 0 {
		slice2 = append(slice2, points.y%10)
		points.y /= 10
	}

	sl1 := []int{}
	sl2 := []int{}

	for i := len(slice) - 1; i >= 0; i-- {
		sl1 = append(sl1, slice[i])
	}

	for i := len(slice2) - 1; i >= 0; i-- {
		sl2 = append(sl2, slice2[i])
	}

	x := "x = "
	y := "y = "

	for _, v := range x {
		z01.PrintRune(v)
	}

	for i := 0; i < len(sl1); i++ {
		z01.PrintRune(rune(sl1[i] + '0'))
	}

	z01.PrintRune(',')
	z01.PrintRune(' ')

	for _, v := range y {
		z01.PrintRune(v)
	}

	for i := 0; i < len(sl2); i++ {
		z01.PrintRune(rune(sl2[i] + '0'))
	}

	z01.PrintRune('\n')
}
