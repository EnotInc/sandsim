package simulation

import (
	"bytes"
	"fmt"
	"time"

	"github.com/enotinc/sandsim/internal/ascii"
)

func (env *Environment) beginRender() {
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()
	fmt.Print(ascii.SaveTerm)
	fmt.Print(ascii.HideCursor)

	for range ticker.C {
		env.render()
	}
}

func (env *Environment) endRender() {
	fmt.Print(
		ascii.ClearAll,
		ascii.Reset,
		ascii.RestoreTerm,
		ascii.ShowCursor)
}

func (env *Environment) render() {
	var output bytes.Buffer
	output.WriteString(ascii.ClearAll)
	output.WriteString(ascii.Reset)

	for i, p := range env.grid {
		if i%env.w == 0 {
			output.WriteString("\n\r")
		}

		output.WriteString(p.String())
	}

	fmt.Print(output.String())
}
