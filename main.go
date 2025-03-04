package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	// go 에는 exception 이 없음 -> 따로 처리해줘야함
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(rw, homeData{PageTitle: "Home"})
}

func main() {
	http.HandleFunc("/", home)

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // exit 1로 종료하는 에러가 있으면 출력

	//chain := blockchain.GetBlockchain()
	//chain.AddBlock("Second Block")
	//chain.AddBlock("Third Block")
	//
	//for _, block := range chain.AllBlocks() {
	//	fmt.Printf("Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %s\n", block.Hash)
	//	fmt.Printf("PrevHash: %s\n\n", block.PrevHash)
	//}
}
