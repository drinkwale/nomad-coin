package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/drinkwale/nomad-coin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tml := template.Must(template.ParseFiles("templates/home.html"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tml.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
