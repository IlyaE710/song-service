package repository

import "github.com/IlyaE710/song-service/internal/v1/song/entity"

type SongRepository interface {
	Save(song entity.Song) (int, error)
}
