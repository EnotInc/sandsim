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

			switch releaseChar {
			case 'm':
				env.spawn = -1
			default:
				env.spawn = x + y*env.w
			}
		}
	}
}
