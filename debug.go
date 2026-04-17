package golib

import (
	"fmt"
	"runtime"
)

func GetStack() []byte {
	stack := make([]byte, 1024)
	n := runtime.Stack(stack, false)
	return stack[:n]
}

func GetCoroutineID() int {
	var id int
	fmt.Sscanf(string(GetStack()), "goroutine %d", &id)
	return id
}
