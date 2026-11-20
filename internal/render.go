package internal

import (
	"bytes"
	"fmt"
)

const (
	reset    = "\033[H"
	clearall = "\033[2J\033[3J"
)

func (env *Environment) render() {
	var output bytes.Buffer
	output.WriteString(clearall)
	output.WriteString(reset)

	for i, p := range env.grid {
		if i%env.w == 0 {
			output.WriteString("\n\r")
		}

		output.WriteString(p.String())
	}

	fmt.Print(output.String())
}
