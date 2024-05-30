package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rwtnitin/gitstore/internal/config"
	"github.com/rwtnitin/gitstore/web"
	"github.com/rwtnitin/gitstore/web/views/components"
	"github.com/rwtnitin/gitstore/web/views/pages"
)

type ApiServer struct {
	cfg config.Config
	// connect to database
	// db     *sql.DB
}

var count int

func NewApiServer(cfg config.Config) *ApiServer {

	return &ApiServer{
		cfg: cfg,
	}
}

func (srv *ApiServer) Run() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pages.HomePage().Render(r.Context(), w)
	})

	router.Get("/count", func(w http.ResponseWriter, r *http.Request) {
		count++
		components.Counter(count).Render(r.Context(), w)
	})

	webAppHandler, err := web.WebAppHandler()
	if err != nil {
		log.Fatalf("unable to set up HTTP server for embedded web app: %v\n", err)
	}
	router.Handle("/static/*", http.StripPrefix("/static", webAppHandler))
	server := &http.Server{
		Addr:    net.JoinHostPort(srv.cfg.HOSTNAME, srv.cfg.Port),
		Handler: router,
	}

	log.Printf("Listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}
}
