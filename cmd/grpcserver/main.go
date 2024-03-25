package main

import (
	"context"
	"io"
	"log"
	"log/slog"
	"net"
	"os"

	"entgo.io/ent/dialect"
	"github.com/ecshreve/backend-playground/ent"
	"github.com/ecshreve/backend-playground/ent/proto/entpb"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func main() {
	// Open or create the log file
	logFile, err := os.OpenFile("logs/grpcserver.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	// Create a multi-writer to write to both standard output and the log file
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	logger := slog.New(slog.NewTextHandler(multiWriter, nil))

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall, logging.PayloadReceived, logging.PayloadSent),
	}

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

	// Run the migration tool (creating tables, etc).
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(InterceptorLogger(logger), opts...),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(InterceptorLogger(logger), opts...),
		),
	)

	// Initialize the generated User service.
	svc := entpb.NewUserService(client)

	// Register the User service with the server.
	entpb.RegisterUserServiceServer(server, svc)

	// Open port 5000 for listening to traffic.
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	// Listen for traffic indefinitely.
	logger.Info("server started", "port", 5000)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
