package simulation

import (
	"os"

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

	quit chan bool
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

		quit:  make(chan bool),
		spawn: 0,
	}
}

func (env *Environment) exit() {
	term.Restore(env.fdin, env.state)
	env.endRender()
	os.Exit(0)
}

func (env *Environment) Run() {
	go env.beginSim()
	go env.beginRender()
	go env.readInput()

	<-env.quit
	env.exit()
}
