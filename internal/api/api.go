package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rocketseat-education/semana-tech-go-react-server/internal/store/pgstore"
)

type apiHandler struct {
	queries *pgstore.Queries
	router  *chi.Mux
}

func (handler apiHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler.router.ServeHTTP(writer, request)
}

func NewHandler(queries *pgstore.Queries) http.Handler {
	api := apiHandler{queries: queries}

	router := chi.NewRouter()
	router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	router.Route("/api", func(routerApi chi.Router) {
		routerApi.Route("/rooms", func(routerRooms chi.Router) {
			routerRooms.Post("/", api.handleCreateRoom)
			routerRooms.Get("/", api.handleGetRooms)
		})
	})

	api.router = router
	return api
}

func (handler apiHandler) handleCreateRoom(writter http.ResponseWriter, request *http.Request) {}
func (handler apiHandler) handleGetRooms(writter http.ResponseWriter, request *http.Request)   {}
