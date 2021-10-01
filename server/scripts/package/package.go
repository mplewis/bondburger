// Package all the plots in the given directory into a gzip file.

package main

import (
	"log"

	"github.com/mplewis/bondburger/server/pack"
	"github.com/mplewis/bondburger/server/util"
)

func main() {
	err := pack.Films(util.MustEnv("PLOTS_DIR"), util.MustEnv("PACKAGE_FILE"))
	if err != nil {
		log.Fatal(err)
	}
}
