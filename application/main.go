package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler).Methods(http.MethodGet)
	r.HandleFunc("/hello", helloHandler).Methods(http.MethodGet) // /hello エンドポイントを追加

	http.ListenAndServe(":8000", r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Root Handler")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "🐰Hello World🐰")
}
