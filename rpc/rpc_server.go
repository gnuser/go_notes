package main

import (
	"errors"
	"net/rpc"
	"net"
	"log"
	"fmt"
	. "./share"
)

type MyInterface struct {}

func (t *MyInterface) Multiply(args *Args, reply *int) error {
	fmt.Printf("multiply with args: %v\n", args)
	*reply = args.A * args.B
	return nil
}

func (t *MyInterface) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(MyInterface)
	rpc.Register(arith)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error: ", e)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}
}