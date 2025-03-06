package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mjcoin/blockchain"
	"mjcoin/utils"
	"net/http"
	"strconv"
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

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
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
			URL:         url("/blocks/{height}"),
			Method:      "GET",
			Description: "See A Block",
		},
	}
	//b, err := json.Marshal(data)
	//utils.HandleErr(err)
	//fmt.Fprintf(rw, "%s", b)

	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(req.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)

	block, err := blockchain.GetBlockchain().GetBlock(id)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrBlockNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)

	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")

	fmt.Printf("Listening on port http://localhost:%d\n", aPort)
	log.Fatal(http.ListenAndServe(port, router))
}
