package services

import (
	"github.com/guiromao/hot-tunes/domain"
	"github.com/guiromao/hot-tunes/errors"
)

type SearchService interface {
	SearchArtist(artistName string) ([]domain.Artist, *errors.AppError)
	SearchNews(artistId string) ([]domain.Album, *errors.AppError)
}
