package render

import (
	"bytes"

	"github.com/enotinc/sandsim/internal/ascii"
)

func Render(line string) string {
	var out bytes.Buffer

	pos := 0
	for pos < len(line) {

		start := pos
		for pos < len(line) && line[pos] == line[start] {
			pos += 1
		}

		paint(&out, line[start:pos])
	}

	out.WriteString("\n\r")

	return out.String()
}

func paint(out *bytes.Buffer, line string) {
	switch line[0] {
	case '*':
		out.WriteString(ascii.Yellow)
	case '#':
		out.WriteString(ascii.Black)
	}

	out.WriteString(line)
	out.WriteString(ascii.Clear)
}
