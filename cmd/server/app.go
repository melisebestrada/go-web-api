package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/melisebestrada/go-web-api/internal/handler"
	"github.com/melisebestrada/go-web-api/internal/repository"
	"github.com/melisebestrada/go-web-api/internal/service"
)

type ConfigServer struct {
	ServerAddress string
	DataFilePath  string
}

type Server struct {
	ServerAddress string
	DataFilePath  string
}

func NewServer(config *ConfigServer) *Server {
	defaultConfig := &ConfigServer{
		ServerAddress: ":8080",
	}

	if config != nil {
		if config.ServerAddress != "" {
			defaultConfig.ServerAddress = config.ServerAddress
		}
		if config.DataFilePath != "" {
			defaultConfig.DataFilePath = config.DataFilePath
		}
	}

	return &Server{
		ServerAddress: defaultConfig.ServerAddress,
		DataFilePath:  defaultConfig.DataFilePath,
	}
}

func (s *Server) Run() (err error) {
	router := chi.NewRouter()

	productRepository, err := repository.NewProductRepository(s.DataFilePath)
	if err != nil {
		return
	}

	productService := service.NewProductService(productRepository)

	productHandler := handler.NewProductHandler(productService)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Route("/products", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/", productHandler.GetAllProducts())
			r.Get("/{id}", productHandler.GetProductById())
			r.Get("/search", productHandler.SearchPriceGt())
			r.Post("/", productHandler.CreateProduct())
			r.Put("/{id}", productHandler.UpdateProduct())
			r.Patch("/{id}", productHandler.PatchProduct())
		})
	})

	err = http.ListenAndServe(s.ServerAddress, router)
	return
}
