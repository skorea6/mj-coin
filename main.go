package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mjcoin/blockchain"
	"mjcoin/utils"
	"net/http"
)

const (
	port string = ":4000"
)

type URL string

// json으로 변환할때 go에서는 marshalText를 사용함. 상속하려면..
func (u URL) MarshalText() (text []byte, err error) { // 시그니처도 동일해야함
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescriptions struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // 데이터가 없으면 없애줌
}

type AddBlockBody struct {
	Message string
}

func (u URLDescriptions) String() string { // interface 코드 구현 (자동 메소드 호출)
	return "Hello I'm url Descriptions!"
}

func documentation(rw http.ResponseWriter, req *http.Request) {
	data := []URLDescriptions{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	//b, err := json.Marshal(data)
	//utils.HandleErr(err)
	//fmt.Fprintf(rw, "%s", b)

	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleErr(json.NewDecoder(req.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

	//explorer.Start()
}
