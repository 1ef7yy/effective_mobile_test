package domain

import "github.com/1ef7yy/effective_mobile_test/internal/models"

func (d *domain) GetSongs(limit, offset int) ([]models.Song, error) {
	songs, err := d.pg.GetSongs(limit, offset)
	if err != nil {
		d.log.Error("error getting songs: " + err.Error())
		return nil, err
	}

	return songs, nil
}
