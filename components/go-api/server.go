package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func newServer(addr string) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello, World! from my Go App",
		})
	})

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

// Bloqueia até receber SIGINT/SIGTERM, encerra o servidor e finaliza o
func gracefulShutdown(server *http.Server) {
	logger := GetLogger()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	received := <-signals

	logger.Info(LogInput{
		Action:  "server",
		Message: fmt.Sprintf("Recebido sinal %s. Encerrando o servidor...", received),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error(LogInput{
			Action:  "server",
			Message: "Erro ao encerrar o servidor.",
			Data:    map[string]any{"error": err.Error()},
		})
		os.Exit(1)
	}

	logger.Info(LogInput{
		Action:  "server",
		Message: "Servidor encerrado com sucesso.",
	})
	os.Exit(1)
}
