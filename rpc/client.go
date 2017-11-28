package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

//服务端存在类型，客户端也必须存在相同类型。

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {

	if len(os.Args) != 2 {
		fmt.println("Usage:", os.Args[0], "server")
	}
	serverAddr := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddr+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Args(17, 8)
	var reply int
	err = client.Call("Math.Multiply", args, reply)
	if err != nil {
		log.Fatal("Match error:", err)
	}

	fmt.Fprintf("Math: %d*%d=%d\n", args.A, args.B, reply)

	var quo Quotient
	err = client.Call("Math.Divide", agrs, &quo)
	if err != nil {
		log.Fatal("Match error:", err)
	}
	fmt.Fprintf("Math: %d/%d=%d remainder %d\n", args.A, args.B, quo.Quo, quo.Rem)
}
