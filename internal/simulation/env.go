package simulation

import (
	"fmt"
	"os"
	"sync"

	"github.com/enotinc/sandsim/internal/ascii"
	"golang.org/x/term"
)

type Environment struct {
	// term window size
	w, h int

	grid []piece

	// additional terminal data
	state *term.State
	fdin  int

	// absolute position. Used to spawn a new piece on it
	spawn int

	// simulation update rate
	rate int

	quit chan bool
	mu   sync.RWMutex
}

func setup() {
	fmt.Print(
		ascii.SaveTerm,
		ascii.HideCursor,
		ascii.EnableMouse,
		ascii.EnableSGR,
	)
}

func reset() {
	fmt.Print(
		ascii.ClearAll,
		ascii.Reset,
		ascii.RestoreTerm,
		ascii.ShowCursor,
		ascii.DisnableMouse,
		ascii.DisnableSGR,
	)
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

	// leaving one line for header
	h -= 1

	_empty := make([]piece, w*h)
	for i := range w * h {
		_empty[i] = empty
	}

	setup()
	return &Environment{
		w: w,
		h: h,

		grid: _empty,

		state: state,
		fdin:  fdin,

		// NOTE: yeah, some kind of magic nuber
		// later this could be moved to some sort of config
		rate: 15,

		spawn: NotFill,

		quit: make(chan bool),
	}
}

func (env *Environment) exit() {
	term.Restore(env.fdin, env.state)
	reset()
	os.Exit(0)
}

func (env *Environment) Run() {
	go env.beginSim()
	go env.beginRender()
	go env.beginReadInput()

	<-env.quit
	env.exit()
}
