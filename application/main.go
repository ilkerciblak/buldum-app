package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("EVERYTHING STARTS HERE!\n")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")
		if _, err := w.Write([]byte(`{"message" : "IT IS ALIVE!}`)); err != nil {
			panic(err)
		}
	})

	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
