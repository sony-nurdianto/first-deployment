package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sony-nurdianto/ecommerce/product_svc/internal/pbgen"
	"github.com/sony-nurdianto/ecommerce/product_svc/internal/service"
	"github.com/sony-nurdianto/ecommerce/product_svc/internal/storage"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	pgi := storage.NewPgInstance()
	dbAddr := os.Getenv("DB_URI")

	db, err := storage.OpenPostgres(dbAddr, pgi)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Database().Close()

	s := grpc.NewServer()
	pbgen.RegisterProductServiceServer(s, &service.ProductService{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	defer lis.Close()

	go func(listener net.Listener) {
		if err := s.Serve(listener); err != nil {
			log.Fatalln(err)
		}
	}(lis)

	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer done()

	var once sync.Once

	for {
		select {
		case <-ctx.Done():
			log.Println("ProductService Server is Shutting Down. Gracefully shutdown")
			s.GracefulStop()
			log.Println("Application Stop.")
			return
		default:
			once.Do(func() {
				log.Println("ProductService Server is Running on 0.0.0.0:50051")
			})
		}
	}
}
