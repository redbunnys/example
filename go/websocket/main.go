package main

import (
	"flag"
	"go-websocket/chat"
	"go-websocket/gorilla"
	"net/http"
)

var (
	addr = flag.String("addr", "localhost:8080", "http service address")
)

func serverHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func serverChat(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/chat" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "chat.html")
}

func main() {
	flag.Parse()
	hub := chat.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		gorilla.GorillaWs(w, r)
	})
	http.HandleFunc("/wschat", func(w http.ResponseWriter, r *http.Request) {
		chat.ChatWs(w, r, hub)
	})
	http.HandleFunc("/chat", serverChat)
	http.HandleFunc("/", serverHome)
	http.ListenAndServe(*addr, nil)
}
