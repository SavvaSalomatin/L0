package http_server

import (
	"html/template"
	"log"
	"net/http"
	"servicel0/config"
)

type Httpsrv struct{}

func home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./http-server/showall.html",
		"./http-server/base.html",
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	exec(ts, config.C.Items(), w)
}

func showonlyone(w http.ResponseWriter, r *http.Request) {
	result, found := config.C.Get(r.URL.Path[len("/show/"):])
	m := result.(config.Order)
	var templatedatainput config.Templatedata

	if found {
		templatedatainput.Order_uid = m.Order_uid
		templatedatainput.Data = string(m.Mjson)
	}

	files := []string{
		"./http-server/show.html",
		"./http-server/base.html",
	}

	if r.URL.Path[:len("/show/")] != "/show/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	exec(ts, templatedatainput, w)
}

func exec(ts *template.Template, templatedatainput any, w http.ResponseWriter) {
	err := ts.Execute(w, templatedatainput)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (server Httpsrv) Httpstart(port *string) {
	log.Println("HTTP-сервер запущен, веб-страница доступна по порту: " + *port)
	http.HandleFunc("/", home)
	http.HandleFunc("/show/", showonlyone)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
