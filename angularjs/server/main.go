package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func lesson8(w http.ResponseWriter, req *http.Request) {
	log.Printf("req:%s", req.FormValue("name"))

	io.WriteString(w, strings.ToUpper(req.Method+"=>"+req.FormValue("name")))
}

func main() {
	http.HandleFunc("/index.php", lesson8)
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
