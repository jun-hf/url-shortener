package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jun-hf/url-shortener/handler"
)

func createDefaultMux() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome!!")
	})
	return mux
}

func main() {
	mux := createDefaultMux()
	
	pathsToUrls := map[string]string{
		"/github": "https://github.com/jun-hf",
		"/url-shortener": "https://github.com/jun-hf/url-shortener",
	}

	mapHandler := handler.MapHandler(pathsToUrls, mux)

	yaml := `
	- path: /csv-quiz
	url: https://github.com/jun-hf/csv-quiz
	- path: /chatroom
	url: https://github.com/jun-hf/chatroom_go
	` 

	yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	log.Println("Starting server at :8080")
	http.ListenAndServe(":8080", yamlHandler)
}
