package c_args

import (
	"unsafe"
)

func ToGoArgs(argc int, argv unsafe.Pointer) []string {
	var args []string
	if argc < 2 {
		return args
	}

	argsPtr := make([]uintptr, argc)
	for i := 0; i < len(argsPtr); i++ {
		argsPtr[i] = *(*uintptr)(unsafe.Pointer(uintptr(argv) + uintptr(i)*unsafe.Sizeof(argv)))
	}
	for _, v := range argsPtr {
		s := charToString(v)
		args = append(args, s)
	}
	return args
}

func charToString(ptr uintptr) string {
	var b []byte
	for {
		c := *(*uint8)(unsafe.Pointer(ptr))
		if c == 0 {
			break
		}
		b = append(b, c)
		ptr++
	}
	return string(b)
}
