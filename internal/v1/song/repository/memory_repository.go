package repository

import "github.com/IlyaE710/song-service/internal/v1/song/entity"

type SongRepositoryInMemoryImpl struct {
	songs map[int]entity.Song
}

func NewSongRepositoryInMemory() SongRepository {
	return &SongRepositoryInMemoryImpl{make(map[int]entity.Song)}
}

func (r *SongRepositoryInMemoryImpl) Save(song entity.Song) (int, error) {
	ID := len(r.songs) + 1
	song.ID = ID
	r.songs[ID] = song

	return ID, nil
}
