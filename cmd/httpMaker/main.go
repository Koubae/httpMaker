package main

import (
	"github.com/Koubae/httpMaker/internal/app"
)

func main() {
	server.Start() // todo: add error handler!
}

func init() {
	server.Configure()
}
