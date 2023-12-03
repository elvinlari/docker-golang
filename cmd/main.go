package main

import (
	"github.com/elvinlari/docker-golang/cmd/servid"
)

func main() {
	server := servid.NewApp()
	server.Start()
}
