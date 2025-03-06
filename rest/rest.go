package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"mjcoin/blockchain"
	"mjcoin/utils"
	"net/http"
)

var port string

type url string

// json으로 변환할때 go에서는 marshalText를 사용함. 상속하려면..
func (u url) MarshalText() (text []byte, err error) { // 시그니처도 동일해야함
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescriptions struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // 데이터가 없으면 없애줌
}

type addBlockBody struct {
	Message string
}

func (u urlDescriptions) String() string { // interface 코드 구현 (자동 메소드 호출)
	return "Hello I'm url Descriptions!"
}

func documentation(rw http.ResponseWriter, req *http.Request) {
	data := []urlDescriptions{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "Get",
			Description: "See All Blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{id}"),
			Method:      "GET",
			Description: "See A Block",
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
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(req.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	handler := http.NewServeMux()

	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)

	fmt.Printf("Listening on port http://localhost:%d\n", aPort)
	log.Fatal(http.ListenAndServe(port, handler))
}
