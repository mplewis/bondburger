package pack

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mplewis/bondburger/server/split"
	"github.com/mplewis/bondburger/server/t"
)

type FilmDump struct {
	Films map[string]t.Film `json:"films"`
}

func load(dir string) (films []*t.Film, err error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.txt"))
	if err != nil {
		return nil, err
	}

	for _, fn := range files {
		f, err := os.Open(fn)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		raw, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		lines := strings.Split(string(raw), "\n")
		year, err := strconv.ParseInt(lines[1], 10, 0)
		if err != nil {
			return nil, err
		}

		title := lines[0]
		plot := strings.Join(lines[3:], "\n")
		plot = strings.ReplaceAll(plot, title, "*******")
		plotLines := split.Sentences(plot)

		films = append(films, &t.Film{
			Title: title,
			Year:  int(year),
			Actor: lines[2],
			Plot:  plotLines,
		})
	}

	return films, nil
}

func dump(fn string, films []*t.Film) error {
	dump := map[string]t.Film{}
	for _, film := range films {
		dump[film.Slug()] = *film
	}

	b, err := json.Marshal(dump)
	if err != nil {
		return err
	}

	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	w, err := gzip.NewWriterLevel(f, gzip.BestCompression)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.Write(b)
	if err != nil {
		return err
	}

	stat, err := os.Stat(fn)
	if err != nil {
		return err
	}

	before := len(b)
	after := stat.Size()
	ratio := float64(before) / float64(after)
	log.Printf("Wrote %d bytes to %s (uncompressed: %d, ratio: %f)", after, fn, before, ratio)
	return nil
}

func Films(fromPlotDir string, toPlotArchive string) error {
	films, err := load(fromPlotDir)
	if err != nil {
		return err
	}
	err = dump(toPlotArchive, films)
	if err != nil {
		return err
	}
	return nil
}
