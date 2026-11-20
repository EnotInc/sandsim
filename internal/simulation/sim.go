package simulation

import (
	"time"
)

func (env *Environment) beginSim() {
	ticker := time.NewTicker(time.Duration(env.rate) * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		env.mu.Lock()
		env.sim()
		env.mu.Unlock()
	}
}

func (env *Environment) sim() {
	if env.spawn != NotFill {
		env.grid[env.spawn] = sand
	}

	env.moved = map[int]bool{}
	size := env.w * env.h

	for i := env.w*env.h - 1; i >= 0; i -= 1 {
		p := env.grid[i]

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
