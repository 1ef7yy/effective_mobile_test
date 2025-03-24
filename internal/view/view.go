package view

import (
	"github.com/1ef7yy/effective_mobile_test/internal/domain"
	"github.com/1ef7yy/effective_mobile_test/pkg/logger"
)

type view struct {
	log    logger.Logger
	domain domain.Domain
}

type View interface {
}

func NewView(logger logger.Logger) (View, error) {
	domain, err := domain.NewDomain(logger)
	if err != nil {
		logger.Error("error initializing domain: " + err.Error())
		return nil, err
	}

	return &view{
		log:    logger,
		domain: domain,
	}, nil
}
