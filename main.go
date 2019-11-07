// main
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func htmlHandler(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open("html/test.html")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		log.Println(err)
	}
}

func streamHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	folder := vars["folder"]
	title := vars["title"]
	path := filepath.Join(folder, title)

	stat, err := os.Stat(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	if stat.IsDir() {
		return
	}

	http.ServeFile(w, r, path)
}

func main() {

	router := mux.NewRouter()

	router.Path("/").Methods(http.MethodGet).HandlerFunc(htmlHandler)
	router.Path("/player/{folder}/{title}").Methods(http.MethodGet).HandlerFunc(streamHandler)

	fmt.Println("start server - http://localhost:8080/player/mp3/더넛츠 - 잔소리.mp3")
	log.Fatal(http.ListenAndServe(":8080", router))

}
