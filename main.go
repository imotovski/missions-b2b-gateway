package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"missions-b2b-gateway/gen/missions"
	"net/http"
)

type ServerError struct {
	Message string
}

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := missions.RegisterMissionServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", authMiddleware(mux))
}

func authMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var authData = r.Header.Get("Authorization")
		if authData != "Basic 123" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			err := json.NewEncoder(w).Encode(ServerError{Message: "Unauthorized"})
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
			return
		}
		h.ServeHTTP(w, r)
	})
}
