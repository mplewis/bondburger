package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mplewis/bondburger/server/split"
)

var key []byte

type Film struct {
	Title string   `json:"title"`
	Year  int      `json:"year"`
	Actor string   `json:"actor"`
	Plot  []string `json:"plot"`
}

type Question struct {
	PlotStep string `json:"plot_step"`
}

func MustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing mandatory environment variable: %s", key)
	}
	return val
}

func LoadFilms(dir string) ([]Film, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.txt"))
	if err != nil {
		return nil, err
	}

	films := []Film{}
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

		plot := split.Sentences(strings.Join(lines[3:], "\n"))
		films = append(films, Film{
			Title: lines[0],
			Year:  int(year),
			Actor: lines[2],
			Plot:  plot,
		})
	}

	return films, nil
}

func encrypt(plaintext string) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	bytes := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func decrypt(ciphertext string) (string, error) {
	cipherbytes, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}
	nonce, todecode := cipherbytes[:nonceSize], cipherbytes[nonceSize:]
	result, err := gcm.Open(nil, nonce, todecode, nil)
	return string(result), err
}

func init() {
	key = []byte(MustEnv("SECRET_KEY"))
	if len(key) > 32 {
		key = key[:32]
	} else if len(key) > 24 {
		key = key[:24]
	} else if len(key) > 16 {
		key = key[:16]
	} else {
		log.Fatal("Secret key must be 16, 24, or 32 bytes")
	}
}

func main() {
	plotDir := MustEnv("PLOTS")
	films, err := LoadFilms(plotDir)
	if err != nil {
		log.Panic(err)
	}

	for _, film := range films {
		fmt.Printf("%s: %d, %s\n", film.Title, film.Year, film.Actor)
		for _, elem := range film.Plot {
			fmt.Printf("    %s\n", elem)
		}
	}

	// r := mux.NewRouter()
	// r.HandleFunc("/films", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(films)
	// })
	// r.HandleFunc("/question", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")

	// })
	// r.HandleFunc("/answer", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(films)
	// }).Methods("POST")
	// fmt.Println("Listening on port 8080")
	// log.Fatal(http.ListenAndServe(":8080", r))
}
