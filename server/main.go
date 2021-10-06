package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mplewis/bondburger/server/pack"
	"github.com/mplewis/bondburger/server/t"
	"github.com/mplewis/bondburger/server/util"
)

var filmsGZ []byte
var filmsCL int

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func init() {
	// load films
	films, err := pack.Load(util.MustEnv("PLOTS_DIR"))
	check(err)

	// to JSON
	dump := map[string]t.Film{}
	for _, film := range films {
		dump[film.Slug()] = *film
	}
	j, err := json.Marshal(dump)
	check(err)

	// to gzipped bytes
	b := bytes.Buffer{}
	w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	check(err)
	_, err = w.Write(j)
	check(err)
	err = w.Close()
	check(err)
	filmsGZ = b.Bytes()

	before := len(j)
	filmsCL = before
	after := len(filmsGZ)
	ratio := float64(before) / float64(after)
	log.Printf("Films: %d bytes (uncompressed: %d, ratio: %f)", after, filmsCL, ratio)
}

func main() {

	r := mux.NewRouter()
	// TODO: Pass static dir in
	// r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	r.HandleFunc("/films.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		// https://gist.github.com/bryfry/09a650eb8aac0fb76c24#gistcomment-2163949
		w.Header().Del("Content-Length")
		w.Header().Set("Content-Length", fmt.Sprint(filmsCL))
		w.Write(filmsGZ)
	})

	host := util.MaybeEnv("HOST", "")
	port := util.MaybeEnv("PORT", "8080")
	hp := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Listening on %s", hp)
	log.Fatal(http.ListenAndServe(hp, r))
}
