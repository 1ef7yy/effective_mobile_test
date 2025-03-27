package db

import (
	"context"
	"fmt"

	"github.com/1ef7yy/effective_mobile_test/internal/errors"
	"github.com/1ef7yy/effective_mobile_test/internal/models"
)

func (p *Postgres) GetSongs(ctx context.Context, limit, offset int, group, song string) ([]models.Song, error) {
	query := `
		SELECT group_name, song, release_date, text, link
		FROM songs
		WHERE (group_name ILIKE $1 OR $1 IS NULL)
		AND (song ILIKE $2 OR $2 IS NULL)
		ORDER BY group_name, song
		LIMIT $3 OFFSET $4
	`

	groupFilter := "%" + group + "%"
	songFilter := "%" + song + "%"

	val, err := p.DB.Query(ctx, query, groupFilter, songFilter, limit, offset)

	if err != nil {
		p.log.Error("error getting songs: " + err.Error())
		return nil, err
	}

	defer val.Close()

	var songs []models.Song

	for val.Next() {
		var song models.Song
		err = val.Scan(&song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link)
		if err != nil {
			p.log.Error("error scanning into song struct: " + err.Error())
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (p *Postgres) GetSongText(ctx context.Context, group, song string) (string, error) {
	val, err := p.DB.Query(ctx, "SELECT text FROM songs WHERE group_name = $1 AND song = $2", group, song)

	if err != nil {
		p.log.Error("error getting song text from db: " + err.Error())
		return "", err
	}

	defer val.Close()

	var text string

	if !val.Next() {
		return "", errors.SongNotFoundErr
	}
	err = val.Scan(&text)
	if err != nil {
		p.log.Error("error scanning into text: " + err.Error())
		return "", err
	}

	return text, nil
}

func (p *Postgres) DeleteSong(ctx context.Context, group, song string) error {
	_, err := p.DB.Query(ctx, "DELETE FROM songs WHERE group_name = $1 AND song = $2", group, song)

	if err != nil {
		p.log.Error("error deleting a song: " + err.Error())
		return err
	}

	return nil
}

func (p *Postgres) CreateSong(ctx context.Context, song models.Song) (models.Song, error) {
	val, err := p.DB.Query(ctx,
		`INSERT INTO songs(group_name, song, release_date, text, link) VALUES($1, $2, $3, $4, $5) RETURNING group_name, song, release_date, text, link`,
		song.Group, song.Song, song.ReleaseDate, song.Text, song.Link)

	if err != nil {
		p.log.Error("error creating a song: " + err.Error())
		return models.Song{}, err
	}

	defer val.Close()

	var songData models.Song

	if !val.Next() {
		p.log.Warn(fmt.Sprintf("song %v already exists", song))
		return models.Song{}, errors.AlreadyExistsErr
	}

	err = val.Scan(&songData.Group, &songData.Song, &songData.ReleaseDate, &songData.Text, &songData.Link)

	if err != nil {
		p.log.Error("error scanning into models.Song struct: " + err.Error())
		return models.Song{}, err
	}

	return songData, nil
}

func (p *Postgres) EditSong(ctx context.Context, editRequest models.EditSongDTO) (models.Song, error) {
	val, err := p.DB.Query(ctx, "UPDATE songs SET release_date=$1, text=$2, link=$3 WHERE group_name=$4 AND song=$5  RETURNING group_name, song, release_date, text, link", editRequest.ReleaseDate, editRequest.Text, editRequest.Link, editRequest.Group, editRequest.Song)

	if err != nil {
		p.log.Error("error editing a song: " + err.Error())
		return models.Song{}, err
	}

	if !val.Next() {
		p.log.Warn(fmt.Sprintf("song %s - %s was not found", editRequest.Group, editRequest.Song))
		return models.Song{}, errors.SongNotFoundErr
	}

	var updatedSong models.Song

	err = val.Scan(&updatedSong.Group, &updatedSong.Song, &updatedSong.ReleaseDate, &updatedSong.Text, &updatedSong.Link)

	if err != nil {
		p.log.Error("error editing a song: " + err.Error())
		return models.Song{}, err
	}

	return updatedSong, nil
}
