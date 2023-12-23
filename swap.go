package piscine

func Swap(a *int, b *int) {
	swA := *a
	swB := *b

	*a = swB
	*b = swA
}
