package routes

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/handler"
	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/pbgen"
	"github.com/sony-nurdianto/ecommerce/ecommerce_web/web"
)

func Routes(svc pbgen.ProductServiceClient) *mux.Router {
	router := mux.NewRouter()

	cssFS, err := fs.Sub(web.TemplateFs, "static/css")
	if err != nil {
		log.Fatal(err)
	}

	router.PathPrefix("/css/").Handler(
		http.StripPrefix("/css/", http.FileServer(http.FS(cssFS))),
	)

	RegisterIndexRoutes(router, svc)

	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)

	return router
}
