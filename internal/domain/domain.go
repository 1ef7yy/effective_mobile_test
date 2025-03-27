package domain

import (
	"context"
	"errors"
	"os"

	"github.com/1ef7yy/effective_mobile_test/internal/models"
	"github.com/1ef7yy/effective_mobile_test/internal/storage/db"
	"github.com/1ef7yy/effective_mobile_test/pkg/logger"
)

type domain struct {
	log logger.Logger
	pg  db.Postgres
}

type Domain interface {
	GetSongs(ctx context.Context, limit, offset int, group, song string) ([]models.Song, error)
	GetText(ctx context.Context, group, song string, limit, offset int) (models.TextResponse, error)
	DeleteSong(ctx context.Context, group, song string) error
	CreateSong(context.Context, models.CreateSongDTO) (models.Song, error)
	EditSong(context.Context, models.EditSongDTO) (models.Song, error)
	CallInfoAPI(group, song string) (models.InfoResponse, error)
}

func NewDomain(logger logger.Logger) (Domain, error) {
	postgresConn, ok := os.LookupEnv("POSTGRES_CONN")
	if !ok {
		logger.Error("could not resolve POSTGRES_CONN in env")
		return nil, errors.New("could not resolve POSTGRES_CONN in env")
	}
	pg, err := db.NewPostgres(context.Background(), postgresConn, logger)
	if err != nil {
		logger.Errorf("Unable to create connection to database: %s", err.Error())
		return nil, err
	}
	return &domain{
		log: logger,
		pg:  *pg,
	}, nil

}
