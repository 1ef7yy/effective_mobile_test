package view

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (v *view) GetSongs(w http.ResponseWriter, r *http.Request) {

	var err error
	var limit int
	var offset int

	limitQuery := r.URL.Query().Get("limit")

	if limitQuery == "" {
		limit = 10
	} else {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			v.log.Error("limit is not a number")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	offsetQuery := r.URL.Query().Get("offset")

	if offsetQuery == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(limitQuery)
		if err != nil {
			v.log.Error("offset is not a number")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	songs, err := v.domain.GetSongs(limit, offset)

	if err != nil {
		v.log.Error("error gettings songs: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if songs == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	resp, err := json.Marshal(songs)

	if err != nil {
		v.log.Error("error marshalling songs: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		v.log.Error("error writing to client: " + err.Error())
		return
	}

}
func (v *view) GetText(w http.ResponseWriter, r *http.Request)    {}
func (v *view) DeleteSong(w http.ResponseWriter, r *http.Request) {}
func (v *view) CreateSong(w http.ResponseWriter, r *http.Request) {}
func (v *view) EditSong(w http.ResponseWriter, r *http.Request)   {}
