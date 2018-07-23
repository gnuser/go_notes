package main

import (
	"fmt"
	"net/rpc"
	"os"
	. "./share"
	)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string!")
		return
	}

	CONNECT := arguments[1]
	c, err := rpc.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	args := Args{16, 3}
	var reply int

	err = c.Call("MyInterface.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Multiply): %v\n", reply)

	var r Quotient
	err = c.Call("MyInterface.Divide", args, &r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Divide): %v\n", r)
}
