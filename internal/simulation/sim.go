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
		switch env.spawn_t {
		case eraser:
			env.grid[env.spawn] = empty
		default:
			if env.grid[env.spawn] == empty {
				env.grid[env.spawn] = env.spawn_t
			}
		}
	}

	size := env.w * env.h

	for i := env.w*env.h - 1; i >= 0; i -= 1 {
		p := env.grid[i]

		below := i + env.w

		switch p {
		case sand:

			if below >= size-env.w { // out of bounds
				continue
			}

			// sand will say on solid blocks due to this if statement
			if env.grid[below] == solid {
				continue
			}

			if env.grid[below] == empty {
				env.grid[i] = empty
				env.grid[below] = sand
				continue
			}

			// if there is an empty space on a line below
			if below%env.w != 0 && env.grid[below-1] == empty {
				env.grid[i] = empty
				env.grid[below-1] = sand
				continue
			}

			// if there is an empty space on a line below
			if below%env.w != env.w-1 && env.grid[below+1] == empty {
				env.grid[i] = empty
				env.grid[below+1] = sand
				continue
			}

		case empty:
			continue
		}
	}
}
