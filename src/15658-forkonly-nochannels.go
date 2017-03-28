package main

import (
	"syscall"
)

func run() {
	for {
		syscall.ForkOnlyBSDTest()
	}
}

func main() {
	forkRoutines := 8

	for i := 0; i < forkRoutines; i++ {
		go run()
	}

	done := make(chan struct{})
	<-done
}
