package main

import (
	"api/internal/config"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {

	cfg := config.LoadEnv()

	r := mux.NewRouter()

	r.HandleFunc("/", YourHandler)

	addr := fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)
	log.Fatal(http.ListenAndServe(addr, r))
}
