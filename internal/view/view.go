package view

import (
	"net/http"

	"github.com/1ef7yy/effective_mobile_test/internal/domain"
	"github.com/1ef7yy/effective_mobile_test/pkg/logger"
)

type view struct {
	log    logger.Logger
	domain domain.Domain
}

type View interface {
	GetSongs(w http.ResponseWriter, r *http.Request)
	GetText(w http.ResponseWriter, r *http.Request)
	DeleteSong(w http.ResponseWriter, r *http.Request)
	CreateSong(w http.ResponseWriter, r *http.Request)
	EditSong(w http.ResponseWriter, r *http.Request)
}

func NewView(logger logger.Logger) (View, error) {
	domain, err := domain.NewDomain(logger)
	if err != nil {
		logger.Errorf("error initializing domain: %s", err.Error())
		return nil, err
	}

	return &view{
		log:    logger,
		domain: domain,
	}, nil
}
