package main

import (
	"all_lb/handlers"
	"all_lb/helpers"
	"all_lb/set"
	"flag"
	"log"
	"net/http"
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
		_, err := w.Write([]byte("hosgeldin"))
		helpers.CheckErr(err)
		return
	})

	http.HandleFunc("/sign", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWsSign(hub, w, r)
	})

	http.HandleFunc("/his", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWsHis(hub, w, r)
	})

	http.HandleFunc("/in", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWsIn(hub, w, r)
	})

	http.HandleFunc("/out", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWsOut(hub, w, r)
	})

	http.HandleFunc("/in-new", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWsInNew(hub, w, r)
	})

	http.HandleFunc("/out-new", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWsOutNew(hub, w, r)
	})

	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
