package simulation

import (
	"fmt"
	"os"
)

func (env *Environment) readInput() {
	buffer := make([]byte, 64)
	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			continue
		}

		input := buffer[:n]

		if n == 1 && input[0] == 'q' {
			env.quit <- true
			return
		}

		if n >= 6 && input[0] == '\033' && input[1] == '[' && input[2] == '<' { // mouse
			var button, x, y int
			var releaseChar rune

			_, err := fmt.Sscanf(string(input[3:]), "%d;%d;%d%c", &button, &x, &y, &releaseChar)
			if err != nil {
				continue
			}

			// fixing mouse ofset in terminal
			// for some reason starting point is 1:1
			if x > 0 && y > 0 { // guard just in case
				x -= 1
				y -= 1
			}

			switch releaseChar {
			case 'm':
				env.spawn = NotFill
			default:
				env.spawn = x + y*env.w
			}
		}
	}
}
