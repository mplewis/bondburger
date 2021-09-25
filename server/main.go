package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mplewis/bondburger/server/split"
)

var key []byte

type Film struct {
	Title string   `json:"title"`
	Year  int      `json:"year"`
	Actor string   `json:"actor"`
	Plot  []string `json:"plot"`
}

func (f *Film) RandomPlot(surround int) []string {
	i := rand.Intn(len(f.Plot)-surround) + surround
	return f.Plot[i-surround : i+surround+1]
}

func (f *Film) Slug() (slug string) {
	slug = strings.ToLower(f.Title)
	slug = regexp.MustCompile(`[^\w\s]`).ReplaceAllLiteralString(slug, "")
	slug = regexp.MustCompile(`\s+`).ReplaceAllLiteralString(slug, "_")
	return
}

type Question struct {
	Plot            []string          `json:"plot"`
	Choices         map[string]string `json:"choices"`
	EncryptedAnswer string            `json:"encrypted_answer"`
}

func GenQuestion(fs []*Film) (*Question, error) {
	const nChoices = 4

	filmInts := make([]int, len(fs))
	for i := range filmInts {
		filmInts[i] = i
	}
	rand.Shuffle(len(filmInts), func(i, j int) {
		filmInts[i], filmInts[j] = filmInts[j], filmInts[i]
	})

	films := []*Film{}
	for i := 0; i < nChoices; i++ {
		films = append(films, fs[filmInts[i]])
	}

	correctInt := rand.Intn(nChoices)
	correctChar := 'A' + correctInt
	fmt.Printf("Correct answer is %s\n", string(byte(correctChar)))

	plot := films[correctInt].RandomPlot(1)

	choices := map[string]string{}
	for i, film := range films {
		filmChar := 'A' + i
		choices[string(byte(filmChar))] = fmt.Sprintf("%s (%s, %d)", film.Title, film.Actor, film.Year)
	}

	encryptedAnswer, err := Encrypt(string(byte(correctChar)))
	if err != nil {
		return nil, err
	}

	return &Question{
		Plot:            plot,
		Choices:         choices,
		EncryptedAnswer: encryptedAnswer,
	}, nil
}

func MustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing mandatory environment variable: %s", key)
	}
	return val
}

func LoadFilms(dir string) (films []*Film, err error) {
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

		films = append(films, &Film{
			Title: title,
			Year:  int(year),
			Actor: lines[2],
			Plot:  plotLines,
		})
	}

	return films, nil
}

type FilmDump struct {
	Films map[string]Film `json:"films"`
}

func DumpFilms(fn string, films []*Film) error {
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	dump := map[string]Film{}
	for _, film := range films {
		dump[film.Slug()] = *film
	}

	return json.NewEncoder(f).Encode(dump)
}

func Encrypt(plaintext string) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(crand.Reader, nonce); err != nil {
		return "", err
	}
	bytes := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func Decrypt(ciphertext string) (string, error) {
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
	rand.Seed(time.Now().UnixNano())
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
	err = DumpFilms("films.json", films)
	if err != nil {
		log.Panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/question", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		q, err := GenQuestion(films)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(q)
	})

	// r.HandleFunc("/answer", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(films)
	// }).Methods("POST")

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
