package main

import "github.com/enotinc/sandsim/internal"

func main() {
	env := internal.Init()
	env.Run()
}
