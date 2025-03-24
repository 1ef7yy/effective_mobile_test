package models

import "time"

type Song struct {
	ReleaseDate time.Time `json:"release_date,omitempty"`
	Text        string    `json:"text,omitempty"`
	Link        string    `json:"link,omitempty"`
	Group       string    `json:"group,omitempty"`
	Song        string    `json:"song,omitempty"`
}
