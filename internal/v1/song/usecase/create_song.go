package usecase

import (
	"github.com/IlyaE710/song-service/internal/v1/song/entity"
	"github.com/IlyaE710/song-service/internal/v1/song/external/query"
	"github.com/IlyaE710/song-service/internal/v1/song/repository"
)

type CreateSongUseCase interface {
	Handle(group string, song string) (*entity.Song, error)
}

type CreateSongUseCaseImpl struct {
	repository repository.SongRepository
}

func NewCreateSongUseCase(repository repository.SongRepository) CreateSongUseCase {
	return &CreateSongUseCaseImpl{repository: repository}
}

func (uc *CreateSongUseCaseImpl) Handle(group string, songTitle string) (*entity.Song, error) {
	q := query.NewGetSongDetailQuery(group, songTitle)
	songDetail, err := q.Execute()
	if err != nil {
		return nil, err
	}
	song := entity.Song{ID: 0, Group: group, Song: songTitle, ReleaseDate: songDetail.ReleaseDate, Text: songDetail.Text, Link: songDetail.Link}
	ID, err := uc.repository.Save(song)
	if err != nil {
		return nil, err
	}
	song.ID = ID
	return &song, nil
}
