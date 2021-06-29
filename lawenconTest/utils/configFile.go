package utils

import (
	"context"
	"encoding/json"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type ConfigModel struct {
	Database struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		Port     string `json:"port"`
		Host     string `json:"host"`
	} `json:"database"`
}

func GetConfigFile() (dataConfig ConfigModel, err error) {
	file, err := os.Open("config.json")
	defer file.Close()

	if err != nil {
		return
	}

	jsonParse := json.NewDecoder(file)
	jsonParse.Decode(&dataConfig)

	return
}

func Start(handler http.Handler) {
	addr := ":8082"

	server := &http.Server{
		Addr:    addr,
		Handler: corsHandler(handler),
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		<-signals

		if err := server.Shutdown(context.Background()); err != nil {

			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {

		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

func corsHandler(handler http.Handler) http.Handler {

	allowOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
	allowHeaders := strings.Split(os.Getenv("CORS_ALLOWED_HEADERS"), ",")
	allowMethods := strings.Split(os.Getenv(\"CORS_ALLOWED_METHODS"), ",")

	crs := cors.New(cors.Options{
		AllowedOrigins: allowOrigins,
		AllowedHeaders: allowHeaders,
		AllowedMethods: allowMethods,
	})
	return crs.Handler(handler)
}
