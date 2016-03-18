package main

import "net/http"

func main() {
	http.HandleFunc("/ws", websocketHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
