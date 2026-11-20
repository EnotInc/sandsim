package simulation

const NotFill = -1

type piece int

const (
	// TODO: extend...
	_ piece = iota
	empty
	sand
)

func (p *piece) String() string {
	switch *p {
	case empty:
		return " "
	case sand:
		return "\033[33m@\033[0m"
	default:
		return ""
	}
}
