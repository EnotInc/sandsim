package ascii

const (
	Reset    = "\033[H"
	ClearAll = "\033[2J\033[3J"

	SaveTerm    = "\033[?1049h"
	RestoreTerm = "\033[?1049l"

	HideCursor = "\033[?25l"
	ShowCursor = "\033[?25h"

	// mouse support
	EnableMouse = "\033[?1002h"
	EnableSGR   = "\033[?1006h"

	DisnableMouse = "\033[?1002l"
	DisnableSGR   = "\033[?1006l"
)
