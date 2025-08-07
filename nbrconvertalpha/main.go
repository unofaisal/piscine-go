package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i, j := range args {
		init := 0
		if i != 0 {
			for _, y := range j {
				w := int(y - '0')
				init = init*10 + w

			}
			if 'a'+init-1 <= 'z' || 'A'+init-1 <= 'Z' {
				if args[1] == "--upper" {
					z01.PrintRune(rune('A' + init - 1))
				} else {
					z01.PrintRune(rune('a' + init - 1))
				}
			} else {
				z01.PrintRune(' ')
			}

		}
	}
	z01.PrintRune('\n')
}
