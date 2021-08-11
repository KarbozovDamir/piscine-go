package main

import (
	"os"
)

func main() {
	args := os.Args[1]
	x := 0
	y := 0
	for i := 0; i < len(args); i++ {
		if args[i] == 'L' {
			y--
		}
		if args[i] == 'R' {
			y++
		}
		if args[i] == 'U' {
			x++
		}
		if args[i] == 'D' {
			x--

		}

	}
	if x == 0 && y == 0 {
		os.Stdout.WriteString("true\n")
		// os.Exit(0)
	} else {
		os.Stdout.WriteString("false\n")
	}
}
