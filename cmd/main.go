package main

// @TODO:

import (
	"fmt"
	"os"

	"github.com/mpragliola/sangennaro/internal/engine"
	"github.com/mpragliola/sangennaro/internal/filesystem"
)

func main() {
	if len(os.Args) > 1 {
		command := os.Args[1]

		switch command {
		case "build":
			fs := new(filesystem.FileSystem)
			e := engine.NewEngine(
				engine.NewSameStructureStrategy(*fs),
				*fs,
			)
			e.Build()
		case "serve":
			fmt.Println("ariunimplemented")
		case "watch":
			fmt.Println("I would watch but ...")
		}
	}
}
