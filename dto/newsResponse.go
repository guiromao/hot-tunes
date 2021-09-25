package dto

import (
	"github.com/guiromao/hot-tunes/domain"
)

type NewsResponse struct {
	NewReleases []domain.Album `json:"new_releases"`
}
