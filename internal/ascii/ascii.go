package ascii

const (
	Reset    = "\033[H"
	ClearAll = "\033[2J\033[3J"

	SaveTerm    = "\033[?1049h"
	RestoreTerm = "\033[?1049l"

	HideCursor = "\033[?25l"
	ShowCursor = "\033[?25h"
)
