package main

import (
	"zinx/src/znet"
)

func main() {
	s := znet.NewServer("[Zinx V0.1]")

	s.Serve()
}
