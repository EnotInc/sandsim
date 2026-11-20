package internal

import (
	"math/rand/v2"
	"os"
	"time"

	"golang.org/x/term"
)

type Environment struct {
	// term window size
	w, h int

	// NOTE: for now this is fine:
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
		// TODO: error message
		panic(err)
	}

	w, h, err := term.GetSize(fdout)
	if err != nil {
		// TODO: error message
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

func (env *Environment) sim() {
	change := rand.IntN(100)

	// just a magic number, to not change sand spanw point too often
	if change >= 97 {
		new := rand.IntN(env.w)
		env.spawn = new
	}

	// spawning a new sand piece
	env.grid[env.spawn] = sand

	env.moved = map[int]bool{}
	size := env.w * env.h

	for i, p := range env.grid {
		below := i + env.w

		switch p {
		case sand:

			if below >= size { // out of bounds
				continue
			}

			if env.moved[i] { // ignoring if this piece alrady was moved
				continue
			}

			if env.grid[below] == empty {
				env.grid[i] = empty
				env.grid[below] = sand
				env.moved[below] = true
				continue
			}

			// if there is an empty space on a line below
			if below%env.w != 0 && env.grid[below-1] == empty {
				env.grid[i] = empty
				env.grid[below-1] = sand
				env.moved[below-1] = true
				continue
			}

			// if there is an empty space on a line below
			if below%env.w != env.w-1 && env.grid[below+1] == empty {
				env.grid[i] = empty
				env.grid[below+1] = sand
				env.moved[below+1] = true
				continue
			}

		case empty:
			continue
		}
	}
}
