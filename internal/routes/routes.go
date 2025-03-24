package routes

import (
	"net/http"

	v1 "github.com/1ef7yy/effective_mobile_test/internal/routes/v1"
	"github.com/1ef7yy/effective_mobile_test/internal/view"
)

func InitRouter(view view.View) *http.ServeMux {
	mux := http.NewServeMux()
	v1 := v1.NewRouter(view)

	mux.Handle("/api/v1/", v1.Endpoints())

	return mux
}
