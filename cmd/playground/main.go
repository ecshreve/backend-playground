package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// Start basic HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, playground")
	})

	slog.Info("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
