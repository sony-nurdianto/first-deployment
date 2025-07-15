package routes

import (
	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/handler"
	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/pbgen"
	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/repository"
	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/service"
)

func RegisterIndexRoutes(router *mux.Router, svc pbgen.ProductServiceClient) {
	productRepo := repository.NewProductRepo(svc)
	service := service.NewProductService(productRepo)
	indexHandler := handler.NewIndexHanlder(service)

	router.HandleFunc("/", indexHandler.Index)
}
