package view

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/1ef7yy/effective_mobile_test/internal/errors"
	"github.com/1ef7yy/effective_mobile_test/internal/models"
)

func (v *view) GetSongs(w http.ResponseWriter, r *http.Request) {

	var err error
	// default values
	var (
		limit  = 10
		offset = 0
	)

	query := r.URL.Query()

	limitQuery := query.Get("limit")

	if limitQuery != "" {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			v.log.Error("limit is not a number")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	offsetQuery := query.Get("offset")

	if offsetQuery != "" {
		offset, err = strconv.Atoi(offsetQuery)
		if err != nil {
			v.log.Error("offset is not a number")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if limit < 0 || offset < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if limit == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	group := query.Get("group")
	song := query.Get("song")

	v.log.Debugf("limitQuery: %s, offsetQuery: %s", limitQuery, offsetQuery)

	songs, err := v.domain.GetSongs(r.Context(), limit, offset, group, song)

	if err == errors.OffsetOutOfRangeErr {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		v.log.Errorf("error gettings songs: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if songs == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	resp, err := json.Marshal(songs)

	if err != nil {
		v.log.Errorf("error marshalling songs: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		v.log.Errorf("error writing to client: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (v *view) GetText(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	group := query.Get("group")

	song := query.Get("song")

	if group == "" || song == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limitQuery := query.Get("limit")

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

	offsetQuery := query.Get("offset")

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

	v.log.Debugf("offset: %d, limit: %d", offset, limit)

	if limit < 0 || offset < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	text, err := v.domain.GetText(r.Context(), group, song, limit, offset)

	if err == errors.OffsetOutOfRangeErr {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(errors.OffsetOutOfRangeErr.Error()))
		if err != nil {
			v.log.Errorf("error writing to client: %s", err.Error())
			return
		}
	}

	if err == errors.SongNotFoundErr {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		v.log.Errorf("error getting text: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(text)
	if err != nil {
		v.log.Errorf("error marshalling text: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resp)

	if err != nil {
		v.log.Errorf("error writing response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
func (v *view) DeleteSong(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	group := query.Get("group")
	song := query.Get("song")

	if group == "" || song == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := v.domain.DeleteSong(r.Context(), group, song)

	if err != nil {
		v.log.Errorf("error deleting a song: %s", err.Error())
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
		v.log.Errorf("error decoding JSON body: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !songRequest.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	song, err := v.domain.CreateSong(r.Context(), songRequest)

	if err == errors.AlreadyExistsErr {
		w.WriteHeader(http.StatusConflict)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			v.log.Errorf("error writing to client: %s", err.Error())
			return
		}
		return
	}

	if err != nil {
		v.log.Errorf("error creating a song: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(song)

	if err != nil {
		v.log.Errorf("error marshalling song: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(resp)
	if err != nil {
		v.log.Errorf("error writing to client: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func (v *view) EditSong(w http.ResponseWriter, r *http.Request) {
	var editRequest models.EditSongDTO

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&editRequest)

	if err != nil {
		v.log.Errorf("error decoding JSON body: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if editRequest.Song == "" || editRequest.Group == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	song, err := v.domain.EditSong(r.Context(), editRequest)

	if err == errors.SongNotFoundErr {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		v.log.Errorf("error editing a song: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(song)

	if err != nil {
		v.log.Errorf("error marshalling song: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	_, err = w.Write(resp)

	if err != nil {
		v.log.Errorf("error writing to client: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
