package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
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

	api.router = router
	return api
}
