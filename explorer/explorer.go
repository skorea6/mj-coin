package explorer

import (
	"fmt"
	"html/template"
	"log"
	"mjcoin/blockchain"
	"net/http"
)

const (
	templateDir string = "explorer/templates/"
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

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(port int) {
	handler := http.NewServeMux()

	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)

	fmt.Printf("Listening on port http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler)) // exit 1로 종료하는 에러가 있으면 출력
}
