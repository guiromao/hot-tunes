package dto

import (
	"github.com/guiromao/hot-tunes/domain"
)

type ArtistResponse struct {
	Artists []domain.Artist `json:"artists"`
}
