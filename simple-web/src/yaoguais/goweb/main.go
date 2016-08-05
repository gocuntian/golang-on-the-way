package main

import (
	"log"
	"net/http"
	"yaoguais/goweb/config"
)

func main() {
	for path, callback := range config.Routers {
		http.HandleFunc(path, callback)
	}
	err := http.ListenAndServe(":9900", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
