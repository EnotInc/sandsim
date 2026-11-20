package main

import "github.com/enotinc/sandsim/internal/simulation"

func main() {
	env := simulation.Init()
	env.Run()
}
