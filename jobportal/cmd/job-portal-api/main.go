package main

import (
	"context"
	"fmt"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/database"
	"job-portal-api/internal/handlers"
	"job-portal-api/internal/services"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func main() {

	err := startApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}
	log.Info().Msg("hello this is our app")

}

func startApp() error {
	log.Info().Msg("main : Started : Initializing authentication support")
	privatePEM, err := os.ReadFile(`C:\Users\ORR Training 22\desktop\jobportal\private.pem`)
	if err != nil {
		return fmt.Errorf("reading auth private key %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("parsing auth private key %w", err)
	}

	publicPEM, err := os.ReadFile(`C:\Users\ORR Training 22\desktop\jobportal\pubkey.pem`)
	if err != nil {
		return fmt.Errorf("reading auth public key %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		return fmt.Errorf("parsing auth public key %w", err)
	}

	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("constructing auth %w", err)
	}

	db, err := database.Connection()
	if err != nil {
		return fmt.Errorf("connection failed")

	}

	ms, err := services.NewService(db)
	if err != nil {
		return fmt.Errorf("initialization  failed")

	}

	// Initialize http service
	api := http.Server{
		Addr:         ":8081",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		IdleTimeout:  800 * time.Second,
		Handler:      handlers.API(a, ms),
	}
	serverErrors := make(chan error, 1)
	go func() {
		log.Info().Str("port", api.Addr).Msg("main: API listening")
		serverErrors <- api.ListenAndServe()
	}()
	//shutdown channel intercepts ctrl+c signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error %w", err)
	case sig := <-shutdown:
		log.Info().Msgf("main: Start shutdown %s", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		//Shutdown gracefully shuts down the server without interrupting any active connections.
		//Shutdown works by first closing all open listeners, then closing all idle connections,
		//and then waiting indefinitely for connections to return to idle and then shut down.
		err := api.Shutdown(ctx)
		if err != nil {
			//Close immediately closes all active net.Listeners
			err = api.Close() // forcing shutdown
			return fmt.Errorf("could not stop server gracefully %w", err)
		}

	}
	return nil

}