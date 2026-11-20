package internal

import (
	"bytes"
	"fmt"
	"time"
)

const (
	reset    = "\033[H"
	clearAll = "\033[2J\033[3J"
)

func (env *Environment) beginRender() {
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		env.render()
	}
}

func (env *Environment) render() {
	var output bytes.Buffer
	output.WriteString(clearAll)
	output.WriteString(reset)

	for i, p := range env.grid {
		if i%env.w == 0 {
			output.WriteString("\n\r")
		}

		output.WriteString(p.String())
	}

	fmt.Print(output.String())
}
