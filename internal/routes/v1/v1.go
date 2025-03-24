package v1

import "github.com/1ef7yy/effective_mobile_test/internal/view"

type Router struct {
	View view.View
}

func NewRouter(view view.View) *Router {
	return &Router{
		View: view,
	}
}
