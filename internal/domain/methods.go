package domain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/1ef7yy/effective_mobile_test/internal/errors"
	"github.com/1ef7yy/effective_mobile_test/internal/models"
)

func (d *domain) GetSongs(limit, offset int) ([]models.Song, error) {
	songs, err := d.pg.GetSongs(limit, offset)
	if err != nil {
		d.log.Error("error getting songs: " + err.Error())
		return nil, err
	}

	return songs, nil
}

func (d *domain) GetText(group, song string, limit, offset int) (models.TextResponse, error) {
	text, err := d.pg.GetSongText(group, song)

	if err != nil {
		d.log.Error("error getting song text: " + err.Error())
		return models.TextResponse{}, err
	}

	verses := strings.Split(text, "\n\n")

	d.log.Debug(fmt.Sprintf("verses: %v", verses))
	d.log.Debug(fmt.Sprintf("verses len: %d", len(verses)))

	if offset >= len(verses) {
		return models.TextResponse{}, errors.OffsetOutOfRangeErr
	}

	end := offset + limit
	if end > len(verses) {
		end = len(verses)
	}

	return models.TextResponse{
		Text: verses[offset:end],
	}, nil
}

func (d *domain) DeleteSong(group, song string) error {
	err := d.pg.DeleteSong(group, song)

	if err != nil {
		d.log.Error("error deleting a song: " + err.Error())
		return err
	}

	return nil
}

func (d *domain) CreateSong(songRequest models.CreateSongDTO) (models.Song, error) {

	info, err := d.CallInfoAPI(songRequest.Group, songRequest.Song)

	if err != nil {
		d.log.Error("error calling external API: " + err.Error())
		return models.Song{}, err
	}

	song := models.Song{
		Group:       songRequest.Group,
		Song:        songRequest.Song,
		ReleaseDate: info.ReleaseDate,
		Text:        info.Text,
		Link:        info.Link,
	}

	err = d.pg.CreateSong(song)
	if err != nil {
		d.log.Error("error creating song: " + err.Error())
		return models.Song{}, err
	}

	return song, nil
}

func (d *domain) EditSong(editRequest models.EditSongDTO) (models.Song, error) {
	song, err := d.pg.EditSong(editRequest)

	if err == errors.SongNotFound {
		return models.Song{}, err
	}

	if err != nil {
		d.log.Error("error editing song: " + err.Error())
		return models.Song{}, nil
	}

	return song, nil
}

func (d *domain) CallInfoAPI(group, song string) (models.InfoResponse, error) {
	externalAPIHost, ok := os.LookupEnv("INFO_SERVER_HOST")
	if !ok {
		return models.InfoResponse{}, fmt.Errorf("could not resolve external API host")
	}

	d.log.Debug("info server host: " + externalAPIHost)

	queries := fmt.Sprintf("group=%s&song=%s", group, song)

	// encoding query (i.e. space -> %20)
	queriesEncoded := url.QueryEscape(queries)

	URL := fmt.Sprintf("%s/info?%s", externalAPIHost, queriesEncoded)

	resp, err := http.Get(URL)
	if err != nil {
		d.log.Error("error GETting external API: " + err.Error())
		return models.InfoResponse{}, err
	}

	var info models.InfoResponse

	respData, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(respData, &info)

	d.log.Debug("response data: " + string(respData))
	if err != nil {
		d.log.Error("error unmarshalling external API: " + err.Error())
		return models.InfoResponse{}, err
	}

	return info, nil
}
