package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ecshreve/backend-playground/ent"
	playgen "github.com/ecshreve/backend-playground/gqlserver"

	_ "github.com/lib/pq"
)

// TODO: this might not be necessary if not trying to log formatted req body.
type GQLBlob struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// TODO make this better or remove it
func logit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				slog.Error("Error reading request body", "error", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			// Restore the io.ReadCloser to its original state
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			// Log the raw request body
			clamp := 80
			if len(bodyBytes) < clamp {
				clamp = len(bodyBytes)
			}
			slog.Info("raw graphql", "query", string(bodyBytes)[:clamp])
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	loglogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(loglogger)
	slog.Info("Starting gqlserver...")

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASS")
	dbname := os.Getenv("POSTGRES_DB")

	// TODO: include missing env vars in error message
	// or just do this better
	if host == "" || port == "" || user == "" || pass == "" || dbname == "" {
		slog.Error("Missing required environment variables")
		os.Exit(1)
	}

	// Connect to the database
	postgresConnStr := "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbname + " password=" + pass + " sslmode=disable"
	client, err := ent.Open(dialect.Postgres, postgresConnStr)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Connected to database at", "host", host, "port", port)

	// Run the auto migration tool
	slog.Info("Running schema migration")
	if err := client.Schema.Create(context.Background()); err != nil {
		slog.Error("failed creating schema resources", "error", err)
		os.Exit(1)
	}
	slog.Info("Schema migration complete")

	// Set up simple HTTP handler
	srv := handler.NewDefaultServer(playgen.NewSchema(client))
	http.Handle("/", logit(playground.Handler("backend-playground", "/query")))
	http.Handle("/query", logit(srv)) // Wrap the srv handler with the logRequestBody function

	// Start the server
	slog.Info("Starting gqlserver on port 8087")
	if err := http.ListenAndServe(":8087", nil); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
	slog.Info("Exiting...")
}
