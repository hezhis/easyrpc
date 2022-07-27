package main

import (
	"github.com/hezhis/easyrpc"
	"github.com/hezhis/easyrpc/protocol"
	"log"
	"time"
)

var c = make(chan *protocol.Message, 10)

type Args struct {
	A int
	B int
}

func add(ctx *easyrpc.Context) {
	var args Args
	if err := ctx.Bind(&args); nil != err {
		log.Println(err)
		return
	}
	log.Printf("ret:%d\n", args.A+args.B)
}

func postRpcReq(req *protocol.Message) {
	c <- req
}

func main() {
	server := easyrpc.NewServer()
	server.AddHandler("service", "add", add)

	client := easyrpc.NewClient(postRpcReq)

	a := 1
	b := 1
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			a += 1
			b *= 2
			client.CallGob("service", "add", &Args{A: a, B: b})
		case req := <-c:
			server.DoCall(req)
		}
	}
}
