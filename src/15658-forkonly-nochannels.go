package main

import (
	"runtime"
	"syscall"
)

func run() {
	for {
		syscall.ForkOnlyBSDTest()
	}
}

func main() {
	forkRoutines := runtime.GOMAXPROCS(0)

	for i := 0; i < forkRoutines; i++ {
		go run()
	}

	done := make(chan struct{})
	<-done
}
