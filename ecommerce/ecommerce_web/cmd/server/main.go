package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/client"
	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/routes"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	clientConn := client.NewClientConnGRPC()
	productSvcAddr := os.Getenv("PRODUCT_SVC_ADDR")
	client, err := client.InitClientGRPC(productSvcAddr, clientConn)
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Conn.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	router := routes.Routes(client.Service)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	go func(s *http.Server) {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}(server)

	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer done()

	var once sync.Once

	for {
		select {
		case <-ctx.Done():
			log.Println("Ecommerce Web Server Is Stoping. Gracefully Shutdown")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				log.Println(err)
			}
			return
		default:
			once.Do(func() {
				log.Println("Server Running At 0.0.0.0:8080")
			})
		}
	}
}
