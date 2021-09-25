package domain

type Album struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Artists     []Artist `json:"artists"`
	Type        string   `json:"album_type"`
	ReleaseDate string   `json:"release_date"`
	TotalTracks string   `json:"total_tracks"`
	Images      []Image  `json:"images"`
	Url         string   `json:"uri"`
}

func (a Album) AlbumToString() string {
	return "Name: " + a.Name + " | Type: " + a.Type +
		" | Release Date: " + a.ReleaseDate + " | Total Tracks: " +
		string(a.TotalTracks)
}
