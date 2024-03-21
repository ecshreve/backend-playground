package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/ecshreve/backend-playground/ent"

	_ "github.com/lib/pq"

	"entgo.io/ent/dialect"
)

func main() {
	// Set up a simple HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, playground")
	})

	// Connect to the database
	client, err := ent.Open(dialect.Postgres, "host=devbox port=5432 user=dbuser dbname=proddb password=dbpass sslmode=disable")
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	defer client.Close()
	slog.Info("Connected to database")

	// Run the auto migration tool
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	slog.Info("Schema migration complete")

	// Start the server (note: this is a blocking call)
	slog.Info("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
