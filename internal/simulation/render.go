package simulation

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/enotinc/sandsim/internal/ascii"
	"github.com/enotinc/sandsim/internal/render"
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

	lines := []string{}
	var line bytes.Buffer

	for i, p := range clone {
		line.WriteString(p.Show())

		if (i+1)%env.w == 0 {
			lines = append(lines, line.String())
			line.Reset()
		}
	}

	for _, l := range lines {
		output.WriteString(render.Render(l))
	}

	output.WriteString(env.bar())
	fmt.Print(output.String())
}

func (env *Environment) bar() string {
	info := fmt.Sprintf("===== [c] clear ===== [q] quit ===== [TAB] - change type [%s] ", env.spawn_t.String())
	controls := ""

	border := strings.Repeat("=", max(env.w-len(info)-len(controls), 0))

	return fmt.Sprintf("%s%s%s", info, controls, border)
}
