package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = false
	return true
}

func IsDoorOpen(Door Door) bool {
	PrintStr("is the Door opened ?")
	return Door.state
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	if ptrDoor.state {
		return true
	}
	return false
}

type Door struct {
	state bool
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = true
}

func main() {
	door := &Door{}

	OpenDoor(door)
	IsDoorClose(door)
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
	if door.state == true {
		CloseDoor(door)
	}
}
