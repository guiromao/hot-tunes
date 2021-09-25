package domain

type ItemsAlbumsJson struct {
	Albums []Album `json:"items"`
}

type ItemsArtistsJson struct {
	Items ItemsArtists `json:"artists"`
}
