package usecase

import (
	"github.com/IlyaE710/song-service/internal/v1/song/entity"
	"github.com/IlyaE710/song-service/internal/v1/song/repository"
)

type AllSongsUseCase interface {
	Handle() ([]entity.Song, error)
}

type AllSongsUseCaseUseCaseImpl struct {
	repository repository.SongRepository
}

func NewAllSongsUseCase(repository repository.SongRepository) AllSongsUseCase {
	return &AllSongsUseCaseUseCaseImpl{repository: repository}
}

func (uc *AllSongsUseCaseUseCaseImpl) Handle() ([]entity.Song, error) {
	entities, err := uc.repository.All()
	if err != nil {
		return nil, err
	}
	return entities, nil
}
