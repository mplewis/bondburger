package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mplewis/bondburger/server/util"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	host := util.MaybeEnv("HOST", "")
	port := util.MaybeEnv("PORT", "8080")
	hp := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Listening on %s", hp)
	log.Fatal(http.ListenAndServe(hp, r))
}
