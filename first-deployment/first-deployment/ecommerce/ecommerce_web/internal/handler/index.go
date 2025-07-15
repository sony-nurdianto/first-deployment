package handler

import (
	"html/template"
	"net/http"

	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/service"
	"github.com/sony-nurdianto/ecommerce/ecommerce_web/web"
)

type IndexHandler struct {
	svc service.ProductService
}

func NewIndexHanlder(svc service.ProductService) *IndexHandler {
	return &IndexHandler{svc: svc}
}

func (ih *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFS(
		web.TemplateFs,
		"static/index.html",
		"static/products.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data, err := ih.svc.Products(0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := template.ExecuteTemplate(w, "index", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
