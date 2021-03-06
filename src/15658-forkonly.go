/*
 * Copyright (c) 2016, Derek Marcotte
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 * 1. Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright
 * notice, this list of conditions and the following disclaimer in the
 * documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package main

/* stdlib includes */
import (
	"fmt"
	"syscall"
)

func run(done chan struct{}) {
	// all the good stuff is private to runtime
	syscall.ForkOnlyBSDTest()

	done <- struct{}{}
	return
}

func main() {
	fmt.Println("exec'ing a bunch of goroutines...")

	runs := 0
	// longest number of garbage collections without crash observed to date
	const maxgc = 25236291
	// runs ~= 60/gc
	const maxruns = 2 * 60 * maxgc

	// 8 & 16 are arbitrary
	done := make(chan struct{}, 8)

	for i := 0; i < 8; i++ {
		go run(done)
		runs++
	}

	for {
		select {
		case <-done:
			go run(done)
			runs++
		}

		if runs > maxruns {
			fmt.Println("Reached maxgc iterations, exiting normally.")
			break
		}
	}
}
