package main

import "C"
import (
	"fmt"
	"unsafe"

	"github.com/mox692/c_args"
)

//export RunGo
func RunGo(argc int, argv unsafe.Pointer) {
	args := c_args.ToGoArgs(argc, argv)
	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d]: %s\n", i, args[i])
	}
}

func main() {
}
