package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"cgt.name/pkg/go-mwclient"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var films []string = []string{
	// "Dr._No_(film)",
	// "From_Russia_with_Love_(film)",
	// "Goldfinger_(film)",
	// "Thunderball_(film)",
	// "You_Only_Live_Twice_(film)",
	// "On_Her_Majesty's_Secret_Service_(film)",
	// "Diamonds_Are_Forever_(film)",
	// "Live_and_Let_Die_(film)",
	// "The_Man_with_the_Golden_Gun_(film)",
	// "The_Spy_Who_Loved_Me_(film)",
	// "Moonraker_(film)",
	// "For_Your_Eyes_Only_(film)",
	// "Octopussy",
	// "A_View_to_a_Kill",
	// "The_Living_Daylights",
	// "Licence_to_Kill",
	// "GoldenEye",
	// "Tomorrow_Never_Dies",
	// "The_World_Is_Not_Enough",
	// "Die_Another_Day",
	// "Casino_Royale_(2006_film)",
	// "Quantum_of_Solace",
	// "Skyfall",
	// "Spectre_(2015_film)",
	// "No_Time_to_Die",
}

func Traverse(doc *html.Node) <-chan *html.Node {
	ch := make(chan *html.Node, 1)

	var recurse func(*html.Node)
	recurse = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			ch <- c
			recurse(c)
		}
	}

	go func() {
		recurse(doc)
		close(ch)
	}()
	return ch
}

func PlotNodes(nodes <-chan *html.Node) <-chan *html.Node {
	ch := make(chan *html.Node, 1)

	go func() {
		in := false
		for node := range nodes {
			if in {
				attrs := []string{}
				for _, attr := range node.Attr {
					attrs = append(attrs, fmt.Sprintf("%s=%s", attr.Key, attr.Val))
				}
				fmt.Printf("[%s] %s\n", strings.Join(attrs, " "), strings.TrimSpace(node.Data))
				ch <- node
			}

			headline := false
			id := false
			for _, attr := range node.Attr {
				if attr.Key == "class" && attr.Val == "mw-headline" {
					headline = true
				}
				if attr.Key == "id" && attr.Val == "Plot" {
					id = true
				}
			}
			if headline && id {
				in = true
			} else if headline {
				for range nodes {
				}
				close(ch)
				return
			}
		}
	}()

	return ch
}

func fetch(client *mwclient.Client, page string) (doc *html.Node, err error) {
	params := map[string]string{
		"action":        "parse",
		"page":          page,
		"prop":          "text",
		"formatversion": "2",
	}
	resp, err := client.Get(params)
	if err != nil {
		return nil, err
	}
	parse, err := resp.GetObject("parse")
	if err != nil {
		return nil, err
	}
	raw, err := parse.GetString("text")
	if err != nil {
		return nil, err
	}
	return html.Parse(strings.NewReader(raw))
}

func save(client *mwclient.Client, page string) error {
	log.Printf("Fetching %s\n", page)

	doc, err := fetch(client, page)
	if err != nil {
		return err
	}

	fn := page + ".txt"
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()
	for node := range PlotNodes(Traverse(doc)) {
		if node.DataAtom == atom.Atom(0) { // data-only
			f.Write([]byte(node.Data))
		}
	}

	log.Printf("Written to %s\n", fn)
	return nil
}

func main() {
	w, err := mwclient.New("https://en.wikipedia.org/w/api.php", "github.com/mplewis/bondburger")
	if err != nil {
		log.Panic(err)
	}

	wg := sync.WaitGroup{}
	for _, film := range films {
		wg.Add(1)
		film := film
		go func() {
			err := save(w, film)
			if err != nil {
				log.Printf("Error: %s: %s\n", film, err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
