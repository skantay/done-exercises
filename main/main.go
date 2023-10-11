package main

import "fmt"
import "piscine"

func main() {

	//piscine.Printcomb()
	//piscine.Printcomb2()
	//piscine.Printnbr()
	//piscine.PrintcombN()
	// fmt.Println(piscine.BasicAtoi("12345"))
	// fmt.Println(piscine.BasicAtoi("0000000012345"))
	// fmt.Println(piscine.BasicAtoi("000000"))

	// fmt.Println(piscine.BasicAtoi2("12345"))
	// fmt.Println(piscine.BasicAtoi2("0000000012345"))
	// fmt.Println(piscine.BasicAtoi2("012 345"))
	// fmt.Println(piscine.BasicAtoi2("Hello World!"))

	// fmt.Println(piscine.Atoi("12345"))
	// fmt.Println(piscine.Atoi("0000000012345"))
	// fmt.Println(piscine.Atoi("012 345"))
	// fmt.Println(piscine.Atoi("Hello World!"))
	// fmt.Println(piscine.Atoi("+1234"))
	// fmt.Println(piscine.Atoi("-1234"))
	// fmt.Println(piscine.Atoi("++1234"))
	// fmt.Println(piscine.Atoi("--1234"))

	s := []int{5,4,3,2,1,0}
	piscine.SortIntegerTable(s)
	fmt.Println(s)
}
