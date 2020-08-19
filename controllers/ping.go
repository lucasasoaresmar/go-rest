package controllers

import (
	"fmt"
	"net/http"
)

type Ping struct{}

func (p *Ping) Pong(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong")
}
