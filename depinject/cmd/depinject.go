package main

import (
	"github.com/lhdv/learn_go_with_tests/depinject"
	"net/http"
	"os"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	depinject.Greet(w, "world")
}

func main() {

	depinject.Greet(os.Stdout, "Foobar")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
