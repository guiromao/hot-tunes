package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guiromao/hot-tunes/domain"
	"github.com/guiromao/hot-tunes/dto"
	"github.com/guiromao/hot-tunes/services"
)

type TunesHandler struct {
	service services.SearchService
}

func (t TunesHandler) SearchArtist(w http.ResponseWriter, req *http.Request) {
	var artistReqDto dto.ArtistRequest

	err := json.NewDecoder(req.Body).Decode(&artistReqDto)
	artistName := artistReqDto.ArtistName

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err)
	}

	artists, e := t.service.SearchArtist(artistName)

	var code int
	var data interface{}

	if e != nil {
		code = e.Code
		data = e.Message
	} else {
		code = http.StatusOK
		data = dto.ArtistResponse{Artists: artists}
	}

	writeResponse(w, code, data)
}

func (t TunesHandler) SearchNews(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	artistId := vars["artist_id"]

	albums, err := t.service.SearchNews(artistId)
	fmt.Println("Here: ", albums)

	var code int
	var data interface{}

	if err != nil {
		code = err.Code
		data = err.Message
	} else if len(albums) == 0 {
		code = http.StatusOK
		data = domain.MessageServer{Text: "No new songs from this artist"}
	} else {
		code = http.StatusOK
		data = dto.NewsResponse{News: albums}
	}

	writeResponse(w, code, data)
}

func NewTunesHandler(s services.SearchService) TunesHandler {
	return TunesHandler{service: s}
}
