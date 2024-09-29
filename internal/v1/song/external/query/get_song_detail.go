package query

import "errors"

type SongDetailResponse struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type SongDetail struct {
	ReleaseDate string
	Text        string
	Link        string
}

type GetSongDetailQuery interface {
	Execute() (*SongDetail, error)
}

type GetSongDetailQueryStubImpl struct {
	group string
	song  string
}

func (q GetSongDetailQueryStubImpl) Execute() (*SongDetail, error) {
	if q.group == "Muse" && q.song == "Supermassive Black Hole" {
		return &SongDetail{
			ReleaseDate: "16.07.2006",
			Text:        "Ooh baby, don't you know I suffer?\\nOoh baby, can you hear me moan?\\nYou caught me under false pretenses\\nHow long before you let me go?\\n\\nOoh\\nYou set my soul alight\\nOoh\\nYou set my soul alight",
			Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
		}, nil
	}
	return nil, errors.New("not found")
}

func NewGetSongDetailQuery(group string, song string) GetSongDetailQuery {
	return &GetSongDetailQueryStubImpl{group, song}
}
