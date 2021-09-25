package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/guiromao/hot-tunes/handlers"
	"github.com/guiromao/hot-tunes/services"
	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func Start() {
	router := mux.NewRouter()
	setEnvVars()
	token, err := createConnection()

	if err != nil {

	}

	searchService := services.NewSearchService(token)
	tunesHandler := handlers.NewTunesHandler(searchService)
	//errorHandler := handlers.NewErrorHandler()

	router.Path("/hottunes/api/v1/search").HandlerFunc(tunesHandler.SearchArtist).Methods("GET")
	router.Path("/hottunes/api/v1/newsfrom/{artist_id}").HandlerFunc(tunesHandler.SearchNews).Methods(http.MethodGet)

	address := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func createConnection() (*oauth2.Token, error) {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		return nil, err
	}

	return token, nil
}
