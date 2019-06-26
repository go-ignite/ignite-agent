package main

import (
	"log"

	agent "github.com/go-ignite/ignite-agent"
)

func main() {
	a, err := agent.Init()
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(a.Start())
}
