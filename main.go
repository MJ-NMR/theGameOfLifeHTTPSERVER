package main

import (
	"log"
	"main/routers"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.HandleFunc("/start", routers.StartHandler)
	http.HandleFunc("/refresh", routers.RefreshHandeler)
	http.HandleFunc("/step", routers.StepHandler)
	http.Handle("/", fs)

	log.Println("Server running on http://localhost:9696")
	log.Fatal(http.ListenAndServe("0.0.0.0:9696", nil))
}
