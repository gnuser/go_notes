package main

import (
	"runtime"
	"fmt"
	"time"
)

func main() {
	exit := make(chan struct{})

	go func() {
		defer close(exit)
		defer println("a")

		func() {
			defer func() {
				println("b", recover() == nil)
			}()

			func() {
				println("c")
				runtime.Goexit()
				println("c done.")
			}()

			println("b done.")
		}()
		println("a done.")
	}()

	<- exit
	println("main exit.")

	for i := 0; i < 2; i++ {
		go func(x int) {
			for n := 0; n < 2; n++ {
				fmt.Printf("%c: %d\n", 'a'+x, n)
				time.Sleep(time.Millisecond)
			}
		}(i)
	}

	runtime.Goexit()
	println("main exit again.")
}