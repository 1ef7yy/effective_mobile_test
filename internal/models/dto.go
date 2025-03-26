package models

import (
	"encoding/json"
	"strings"
	"time"
)

type Song struct {
	Group       string    `json:"group,omitempty"`
	Song        string    `json:"song,omitempty"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
	Text        string    `json:"text,omitempty"`
	Link        string    `json:"link,omitempty"`
}

type CreateSongDTO struct {
	Group string `json:"group,omitempty"`
	Song  string `json:"song,omitempty"`
}

type InfoResponse struct {
	ReleaseDate time.Time `json:"releaseDate,omitempty"`
	Text        string    `json:"text,omitempty"`
	Link        string    `json:"link,omitempty"`
}

type EditSongDTO struct {
	Group       string    `json:"group,omitempty"`
	Song        string    `json:"song,omitempty"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
	Text        string    `json:"text,omitempty"`
	Link        string    `json:"link,omitempty"`
}

type TextResponse struct {
	Text []string `json:"text,omitempty"`
}

func (c CreateSongDTO) Valid() bool {
	return c.Group != "" && c.Song != ""
}

func (i *InfoResponse) UnmarshalJSON(data []byte) error {
	var aux struct {
		ReleaseDate string
		Text        string
		Link        string
	}

	err := json.Unmarshal(data, &aux)

	if err != nil {
		return err
	}

	i.ReleaseDate, err = time.Parse("02.01.2006", aux.ReleaseDate)

	if err != nil {
		return err
	}

	// some requests' \n are turned into \\n for escape char purposes
	// so we change all \\n into \n
	i.Text = strings.ReplaceAll(aux.Text, "\\n", "\n")
	i.Link = aux.Link
	return nil
}
