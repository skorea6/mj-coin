package main

import (
	"fmt"
	"html/template"
	"log"
	"mjcoin/blockchain"
	"net/http"
)

const port string = ":4000"

type homeData struct {
	PageTitle string // 공유를 위해서는 대문자
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(rw, homeData{
		PageTitle: "Home",
		Blocks:    blockchain.GetBlockchain().AllBlocks(),
	})
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
