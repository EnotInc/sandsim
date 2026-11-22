package simulation

const NotFill = -1

type piece int

const (
	_ piece = iota

	empty
	sand
	solid
	eraser

	pieceAmount
)

// return the symbol of the piece
func (p *piece) Show() string {
	switch *p {
	case empty:
		return " "
	case sand:
		return "*"
	case solid:
		return "#"
	default:
		return ""
	}
}

// return the name of the piece
func (p *piece) String() string {
	switch *p {
	case empty:
		return "empty"
	case sand:
		return "sand"
	case solid:
		return "solid"
	case eraser:
		return "eraser"
	default:
		return ""
	}
}

func (env *Environment) nextPiece() {
	env.spawn_t += 1

	if env.spawn_t == pieceAmount {
		env.spawn_t = sand
	}
}
