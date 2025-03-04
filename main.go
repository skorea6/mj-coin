package main

import (
	"fmt"
	"html/template"
	"log"
	"mjcoin/blockchain"
	"net/http"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string // 공유를 위해서는 대문자
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(rw, "home", homeData{
		PageTitle: "Home",
		Blocks:    blockchain.GetBlockchain().AllBlocks(),
	})
}

func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	http.HandleFunc("/", home)

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // exit 1로 종료하는 에러가 있으면 출력
}
