package internal

import (
	"os"
	"time"

	"golang.org/x/term"
)

type Environment struct {
	// term window size
	w, h int

	grid  []piece
	moved map[int]bool

	// additional terminal data
	state *term.State
	fdin  int

	spawn int
}

func Init() *Environment {
	fdin := int(os.Stdin.Fd())
	fdout := int(os.Stdout.Fd())

	state, err := term.MakeRaw(fdin)
	if err != nil {
		panic(err)
	}

	w, h, err := term.GetSize(fdout)
	if err != nil {
		panic(err)
	}

	h -= 1

	_empty := make([]piece, w*h)
	for i := range w * h {
		_empty[i] = empty
	}

	return &Environment{
		w: w,
		h: h,

		grid:  _empty,
		moved: make(map[int]bool),

		state: state,
		fdin:  fdin,

		spawn: 0,
	}
}

func (env *Environment) Run() {
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {

		env.sim()
		env.render()
	}

}
