package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune('t')
	z01.PrintRune('e')
	z01.PrintRune('s')
	z01.PrintRune('t')
	var s = "Héllo!"
	var l int = len(s)
	z01.PrintRune(rune(l))
	z01.PrintRune('\n')

	//go run test/main.go

}
