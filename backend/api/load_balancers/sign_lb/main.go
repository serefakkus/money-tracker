package main

import (
	"flag"
	"log"
	"net/http"
	"sign_lb/handlers"
	"sign_lb/set"
)

var addr = flag.String("addr", ":80", "http service address")

func main() {
	set.InitAllStrings()

	flag.Parse()
	hub := handlers.NewHub()
	go hub.Run()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	})

	http.HandleFunc("/sign", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWsSign(hub, w, r)
	})

	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
