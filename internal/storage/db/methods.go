package db

import (
	"context"
	"fmt"

	"github.com/1ef7yy/effective_mobile_test/internal/errors"
	"github.com/1ef7yy/effective_mobile_test/internal/models"
)

func (p *Postgres) GetSongs(limit, offset int) ([]models.Song, error) {
	// context work?
	val, err := p.DB.Query(context.Background(), "SELECT group_name, song, release_date, text, link FROM songs ORDER BY song LIMIT $1 OFFSET $2", limit, offset)

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

func (p *Postgres) GetSongText(group, song string) (string, error) {
	val, err := p.DB.Query(context.Background(), "SELECT text FROM songs WHERE group_name = $1 AND song = $2", group, song)

	if err != nil {
		p.log.Error("error getting song text from db: " + err.Error())
		return "", err
	}

	defer val.Close()

	var text string

	if val.Next() {
		err = val.Scan(&text)
		if err != nil {
			p.log.Error("error scanning into text: " + err.Error())
			return "", err
		}
	}

	return text, nil
}

func (p *Postgres) DeleteSong(group, song string) error {
	_, err := p.DB.Query(context.Background(), "DELETE FROM songs WHERE group_name = $1 AND song = $2", group, song)

	if err != nil {
		p.log.Error("error deleting a song: " + err.Error())
		return err
	}

	return nil
}

func (p *Postgres) CreateSong(song models.Song) error {
	_, err := p.DB.Query(context.Background(), "INSERT INTO songs(group_name, song, release_date, text, link) VALUES($1, $2, $3, $4, $5)", song.Group, song.Song, song.ReleaseDate, song.Text, song.Link)

	if err != nil {
		p.log.Error("error creating a song: " + err.Error())
		return err
	}

	return nil
}

func (p *Postgres) EditSong(editRequest models.EditSongDTO) (models.Song, error) {
	val, err := p.DB.Query(context.Background(), "UPDATE songs SET release_date=$1, text=$2, link=$3 WHERE group_name=$4 AND song=$5 RETURNING group_name, song, release_date, text, link", editRequest.ReleaseDate, editRequest.Text, editRequest.Link, editRequest.Group, editRequest.Song)

	if err != nil {
		p.log.Error("error editing a song: " + err.Error())
		return models.Song{}, err
	}

	if !val.Next() {
		p.log.Warn(fmt.Sprintf("song %s - %s was not found", editRequest.Group, editRequest.Song))
		return models.Song{}, errors.SongNotFound
	}

	var updatedSong models.Song

	err = val.Scan(&updatedSong.Group, &updatedSong.Song, &updatedSong.ReleaseDate, &updatedSong.Text, &updatedSong.Link)

	if err != nil {
		p.log.Error("error editing a song: " + err.Error())
		return models.Song{}, err
	}

	return updatedSong, nil
}
