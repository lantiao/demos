package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func lesson8(w http.ResponseWriter, req *http.Request) {
	log.Printf("req:%s", req.FormValue("name"))

	io.WriteString(w, strings.ToUpper(req.Method+"=>"+req.FormValue("name")))
}

type Data struct {
	Name string `json:"name"`
	Book string `json:"book"`
}

func getOne(w http.ResponseWriter, req *http.Request) {
	list := Data{
		Name: "lantiao",
		Book: "angularjs_" + req.FormValue("name"),
	}
	b, _ := json.Marshal(list)
	io.WriteString(w, string(b))
}

func getArr(w http.ResponseWriter, req *http.Request) {
	list := []Data{
		Data{
			Name: "lantiao",
			Book: "angularjs_" + req.FormValue("name"),
		},
		Data{
			Name: "lantiao2",
			Book: "angularjs2_" + req.FormValue("name"),
		},
	}
	b, _ := json.Marshal(list)
	io.WriteString(w, string(b))
}

func main() {
	http.HandleFunc("/index.php", lesson8)
	http.HandleFunc("/getOne.php", getOne)
	http.HandleFunc("/getArr.php", getArr)
	dir := http.Dir("../")
	log.Printf("dir:%s", dir)
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(dir)))
	port := ":8080"
	log.Printf("ListenAndServe at %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
