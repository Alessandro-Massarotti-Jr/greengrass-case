package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	loadDotEnv()

	logger := GetLogger()

	port := 3001
	if value := os.Getenv("PORT"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			port = parsed
		}
	}

	server := newServer(fmt.Sprintf(":%d", port))

	go func() {
		logger.Info(LogInput{
			Action:  "main",
			Message: "Servidor rodando",
			Data: map[string]any{
				"address": server.Addr,
				"port":    port,
			},
		})

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(LogInput{
				Action:  "main",
				Message: "Erro ao iniciar o servidor.",
				Data:    map[string]any{"error": err.Error()},
			})
			os.Exit(1)
		}
	}()

	gracefulShutdown(server)
}
