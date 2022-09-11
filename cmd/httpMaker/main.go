package main

import (
	"github.com/Koubae/httpMaker/internal/app"
)

func main() {
	server.Start()
}

func init() {
	server.Configure()
}
