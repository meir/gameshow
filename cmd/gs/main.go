package main

import "github.com/meir/gameshow/internal/gs_server"

func main() {
	if err := gs_server.NewServer().Listen(); err != nil {
		panic(err)
	}
}
