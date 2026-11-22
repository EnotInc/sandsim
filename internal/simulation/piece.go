package simulation

const NotFill = -1

type piece int

const (
	_ piece = iota

	empty
	sand
	solid

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
	default:
		return ""
	}
}

func (env *Environment) nextPiece() {
	env.spaw_t += 1

	if env.spaw_t == pieceAmount {
		env.spaw_t = sand
	}
}
