package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Creating a handler

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("hello") //logger has println method
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "nono", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "hmm interesting %s", d)

}
