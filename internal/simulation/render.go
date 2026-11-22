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

		output.WriteString(p.Show())
	}

	output.WriteString(env.bar())
	fmt.Print(output.String())
}

func (env *Environment) bar() string {
	info := fmt.Sprintf("===== [q] quit ===== [TAB] - change type [%s] ", env.spaw_t.String())
	controls := ""

	border := strings.Repeat("=", max(env.w-len(info)-len(controls), 0))

	return fmt.Sprintf("%s%s%s", info, controls, border)
}
