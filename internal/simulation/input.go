package simulation

import "os"

func (env *Environment) readInput() {
	buffer := make([]byte, 64)
	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			break
		}

		input := buffer[:n]

		if input[0] == 'q' {
			env.quit <- true
		}
	}
}
