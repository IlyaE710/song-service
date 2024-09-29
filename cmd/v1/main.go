package main

import (
	"github.com/IlyaE710/song-service/internal/v1/song/repository"
	"github.com/IlyaE710/song-service/internal/v1/song/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createSongRequest struct {
	Group string `json:"group" binding:"required"`
	Song  string `json:"song" binding:"required"`
}

type createSongResponse struct {
	ID          int    `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func main() {
	repo := repository.NewSongRepositoryInMemory()
	r := setupRouter(repo)

	if err := r.Run("127.0.0.1:8081"); err != nil {
		println(err.Error())
	}
}

func setupRouter(repo repository.SongRepository) *gin.Engine {
	r := gin.Default()
	r.POST("/api/v1/song", createSongHandler(repo))
	r.GET("/api/v1/songs", allSongsHandler(repo))
	return r
}

func createSongHandler(repo repository.SongRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request createSongRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			handleError(c, err, http.StatusBadRequest)
			return
		}

		uc := usecase.NewCreateSongUseCase(repo)
		song, err := uc.Handle(request.Group, request.Song)
		if err != nil {
			handleError(c, err, http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Song created successfully",
			"song":    createSongResponse{song.ID, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link},
		})
	}
}

func allSongsHandler(repo repository.SongRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		uc := usecase.NewAllSongsUseCase(repo)
		songs, err := uc.Handle()
		if err != nil {
			handleError(c, err, http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, songs)
	}
}

func handleError(c *gin.Context, err error, status int) {
	c.JSON(status, gin.H{"error": err.Error()})
}
