package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Serve the index page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = fmt.Fprintln(w, "<html><body><h1>Hello, World!</h1></body></html>")
	})

	// Start the server in a goroutine
	server := &http.Server{Addr: ":9090"}
	go func() {
		fmt.Println("Server started on :9090")
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server error:", err)
		}
	}()

	// Wait for 90 seconds and then crash
	time.Sleep(90 * time.Second)
	fmt.Println("Crashing the server after 90 seconds")
	panic("Server crashed intentionally")
}
