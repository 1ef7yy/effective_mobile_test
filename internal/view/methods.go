package view

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/1ef7yy/effective_mobile_test/internal/errors"
	"github.com/1ef7yy/effective_mobile_test/internal/models"
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
		offset, err = strconv.Atoi(offsetQuery)
		if err != nil {
			v.log.Error("offset is not a number")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	v.log.Debug(fmt.Sprintf("limitQuery: %s, offsetQuery: %s", limitQuery, offsetQuery))

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
func (v *view) GetText(w http.ResponseWriter, r *http.Request) {

	group := r.URL.Query().Get("group")

	song := r.URL.Query().Get("song")

	if group == "" || song == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limitQuery := r.URL.Query().Get("limit")

	var (
		limit  int
		offset int
		err    error
	)

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
		offset, err = strconv.Atoi(offsetQuery)
		if err != nil {
			v.log.Error("offset is not a number")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	v.log.Debug(fmt.Sprintf("offset: %d, limit: %d", offset, limit))

	text, err := v.domain.GetText(group, song, limit, offset)

	if err == errors.OffsetOutOfRangeErr {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(errors.OffsetOutOfRangeErr.Error()))
		if err != nil {
			v.log.Error("error writing to client: " + err.Error())
			return
		}
	}

	if err != nil {
		v.log.Error("error getting text: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(text)
	if err != nil {
		v.log.Error("error marshalling text: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)

	if err != nil {
		v.log.Error("error writing response: " + err.Error())
		return
	}

}
func (v *view) DeleteSong(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	song := r.URL.Query().Get("song")

	if group == "" || song == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := v.domain.DeleteSong(group, song)

	if err != nil {
		v.log.Error("error deleting a song: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func (v *view) CreateSong(w http.ResponseWriter, r *http.Request) {
	var songRequest models.CreateSongDTO

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&songRequest)

	if err != nil {
		v.log.Error("error decoding JSON body: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !songRequest.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	song, err := v.domain.CreateSong(songRequest)

	if err != nil {
		v.log.Error("error creating a song: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(song)

	if err != nil {
		v.log.Error("error marshalling song: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		v.log.Error("error writing to client: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
func (v *view) EditSong(w http.ResponseWriter, r *http.Request) {
	var editRequest models.EditSongDTO

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&editRequest)

	if err != nil {
		v.log.Error("error decoding JSON body: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	song, err := v.domain.EditSong(editRequest)

	if err == errors.SongNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		v.log.Error("error editing a song: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(song)

	if err != nil {
		v.log.Error("error marshalling song: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	_, err = w.Write(resp)

	if err != nil {
		v.log.Error("error writing to client: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
