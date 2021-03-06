package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/guiromao/hot-tunes/domain"
	"github.com/guiromao/hot-tunes/errors"
	errs "github.com/guiromao/hot-tunes/errors"
	"golang.org/x/oauth2"
)

type SearchServiceSpotify struct {
	Token *oauth2.Token
}

func (serv SearchServiceSpotify) SearchArtist(artistName string) ([]domain.Artist, *errors.AppError) {
	searchArtist := strings.ReplaceAll(strings.Trim(artistName, " "), " ", "+")
	searchQuery := "https://api.spotify.com/v1/search?query=" + searchArtist + "&offset=0&limit=5&type=artist"

	body, err := serv.doApiCall(searchQuery)

	if err != nil {
		return nil, errors.NewUnexpectedError("Error retrieving artists")
	}

	var itemsArtists domain.ItemsArtistsJson
	json.Unmarshal([]byte(string(body)), &itemsArtists)

	return itemsArtists.Items.Artists, nil
}

func (s SearchServiceSpotify) doApiCall(query string) ([]byte, error) {
	client := &http.Client{}
	bearer := "Bearer " + string(s.Token.AccessToken)
	req, err := http.NewRequest("GET", query, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", bearer)
	resp, err := client.Do(req)

	if err == nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return nil, err
	}

	return body, nil
}

func (s SearchServiceSpotify) SearchNews(artistId string) ([]domain.Album, *errors.AppError) {
	baseUrl := "https://api.spotify.com/v1/artists/" + artistId + "/albums?offset=0&limit=5&include_groups="
	urlAlbums := baseUrl + "album"
	urlSingles := baseUrl + "single"

	albums, err1 := s.getAlbums(urlAlbums)
	singles, err2 := s.getAlbums(urlSingles)

	if err1 != nil || err2 != nil {
		return nil, errs.NewUnexpectedError("Error fetching new albums and singles from artist")
	}

	allAlbums := append(albums, singles...)
	OrderByDate(allAlbums)

	return allAlbums, nil
}

func (s SearchServiceSpotify) getAlbums(url string) ([]domain.Album, *errs.AppError) {
	body, err := s.doApiCall(url)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var items domain.ItemsAlbumsJson
	json.Unmarshal([]byte(string(body)), &items)

	return removeDuplicates(FilterRecentAlbums(items.Albums)), nil
}

func OrderByDate(albums []domain.Album) {
	sort.Slice(albums, func(i, j int) bool {
		firstDate, _ := time.Parse("2006-01-02", albums[i].ReleaseDate)
		secondDate, _ := time.Parse("2006-01-02", albums[j].ReleaseDate)

		return firstDate.After(secondDate)
	})
}

func FilterRecentAlbums(albums []domain.Album) []domain.Album {
	result := []domain.Album{}
	now := time.Now()
	daysThresh := now.AddDate(0, 0, -30)

	for _, album := range albums {
		albumDate, _ := time.Parse("2006-01-02", album.ReleaseDate)

		if albumDate.After(daysThresh) {
			result = append(result, album)
		}
	}

	return result
}

func removeDuplicates(albums []domain.Album) []domain.Album {
	albumMap := make(map[string]domain.Album, 0)

	for _, album := range albums {
		if _, present := albumMap[album.Name]; !present {
			albumMap[album.Name] = album
		}
	}

	return mapToSlice(albumMap)
}

func mapToSlice(albumMap map[string]domain.Album) []domain.Album {
	result := []domain.Album{}

	for _, album := range albumMap {
		result = append(result, album)
	}

	return result
}

func printAlbums(albums []domain.Album) {
	if len(albums) == 0 {
		fmt.Println("No recent albums released.")
	} else {
		for _, album := range albums {
			fmt.Println(album.AlbumToString())
		}
	}
}

func NewSearchService(token *oauth2.Token) SearchServiceSpotify {
	return SearchServiceSpotify{
		Token: token,
	}
}
