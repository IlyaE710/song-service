package main

import (
	"github.com/IlyaE710/song-service/internal/v1/song/repository"
	"github.com/IlyaE710/song-service/internal/v1/song/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	repo := repository.NewSongRepositoryInMemory()

	r.POST("api/v1/song", func(c *gin.Context) {
		type createSongRequest struct {
			Group string `json:"group" binding:"required"`
			Song  string `json:"song" binding:"required"`
		}

		var request createSongRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		uc := usecase.NewCreateSongUseCase(repo)
		song, err := uc.Handle(request.Group, request.Song)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		type createSongResponse struct {
			ID          int
			Group       string
			Song        string
			ReleaseDate string
			Text        string
			Link        string
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Song created successfully",
			"song":    createSongResponse{song.ID, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link},
		})
	})

	err := r.Run("127.0.0.1:8081")
	if err != nil {
		println(err.Error())
	}
}
