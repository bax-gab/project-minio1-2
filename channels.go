/*package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

*/

// run

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that unbuffered channels act as pure fifos.

package main

import "os"

const factorN = 10

func AsynchFifo() {
	ch := make(chan int, factorN)
	for i := 0; i < factorN; i++ {
		ch <- i
	}
	for i := 0; i < factorN; i++ {
		if <-ch != i {
			print("bad receive\n")
			os.Exit(1)
		}
	}
}

func Chain(ch <-chan int, val int, in <-chan int, out chan<- int) {
	<-in
	if <-ch != val {
		panic(val)
	}
	out <- 1
}

// thread together a daisy chain to read the elements in sequence
func SynchFifo() {
	ch := make(chan int)
	in := make(chan int)
	start := in
	for i := 0; i < factorN; i++ {
		out := make(chan int)
		go Chain(ch, i, in, out)
		in = out
	}
	start <- 0
	for i := 0; i < factorN; i++ {
		ch <- i
	}
	<-in
}

func main() {
	AsynchFifo()
	SynchFifo()
}
