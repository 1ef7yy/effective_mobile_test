package db

import (
	"context"

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

func (p *Postgres) DeleteSong(group, song string) error {
	_, err := p.DB.Query(context.Background(), "DELETE FROM songs WHERE group_name = $1 AND song = $2", group, song)

	if err != nil {
		p.log.Error("error deleting a song: " + err.Error())
		return err
	}

	return nil
}
