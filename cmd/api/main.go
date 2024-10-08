package main

import (
	"log/slog"
	"net/http"

	"prime/internal"
)

func main() {
	handlers := internal.InitializeServer()

	port := "3030"
	slog.Info("server starting on", "port", port)
	if err := http.ListenAndServe("localhost:"+port, handlers); err != nil {
		slog.Error("failed to start server", err)
	}
}
