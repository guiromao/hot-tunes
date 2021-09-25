package dto

import (
	"github.com/guiromao/hot-tunes/domain"
)

type NewsResponse struct {
	News []domain.Album `json:"news"`
}
