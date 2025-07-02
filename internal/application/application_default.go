package application

import (
	"app/internal/handler"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ConfigServerChi struct {
	ServerAddress  string
	LoaderFilePath string
}

func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

type ServerChi struct {
	serverAddress  string
	loaderFilePath string
}

func (a *ServerChi) Run() (err error) {
	ld := loader.NewProductJSONFile(a.loaderFilePath)
	db, err := ld.Load()
	if err != nil {
		return
	}
	rp := repository.NewProductMap(db)
	sv := service.NewProductDefault(rp)
	hd := handler.NewProductDefault(sv)
	rt := chi.NewRouter()

	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})

	rt.Route("/products", func(rt chi.Router) {
		rt.Get("/", hd.GetAll())
		rt.Post("/", hd.CreateProducts())
		rt.Get("/search", hd.SearchProducts())
		rt.Get("/{id_product}", hd.GetProductById())
		rt.Put("/{id_product}", hd.UpdateProduct())
		rt.Patch("/{id_product}", hd.UpdateProductAttributes())
		rt.Delete("/{id_product}", hd.DeleteProduct())
	})

	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
