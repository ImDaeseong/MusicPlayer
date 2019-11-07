// Music
package main

import (
	"fmt"
	"net/http"
)

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/test.html")
}

func mp3Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./mp3/더넛츠 - 잔소리.mp3")
	fmt.Println("mp3 playing")
}

func main() {
	http.HandleFunc("/player", mp3Handler)
	http.HandleFunc("/", htmlHandler)

	/*
		http.ListenAndServeTLS(":8080", "https-server.crt", "https-server.key", nil);
		fmt.Println("https://localhost:8080")
	*/

	http.ListenAndServe(":8080", nil)
	fmt.Println("http://localhost:8080")
}
