package simulation

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/enotinc/sandsim/internal/ascii"
)

func (env *Environment) beginRender() {
	ticker := time.NewTicker(time.Duration(env.rate) * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		env.render()
	}
}

func (env *Environment) render() {
	var output bytes.Buffer
	output.WriteString(ascii.Reset)

	env.mu.Lock()
	clone := make([]piece, env.w*env.h)
	copy(clone, env.grid)
	env.mu.Unlock()

	for i, p := range clone {
		if i%env.w == 0 {
			output.WriteString("\n\r")
		}

		output.WriteString(p.String())
	}

	output.WriteString(env.header())
	fmt.Print(output.String())
}

func (env *Environment) header() string {
	title := " press 'q' to quit "
	border := strings.Repeat("=", (env.w-len(title))/2)

	return fmt.Sprintf("%s%s%s", border, title, border)
}
