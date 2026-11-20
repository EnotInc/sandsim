package internal

import (
	"math/rand/v2"
)

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
