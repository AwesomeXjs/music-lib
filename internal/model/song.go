package model

type Song struct {
	Id          string `json:"id" db:"id"`
	Group       string `json:"group" db:"group"`
	Song        string `json:"song" db:"song"`
	ReleaseDate string `json:"release_date" db:"release_date"`
	Text        string `json:"text" db:"text"`
	Patronymic  string `json:"patronymic" db:"patronymic"`
}

type SongCreate struct {
	Group string `json:"group" validate:"required"`
	Song  string `json:"song" validate:"required"`
}

type SongUpdate struct {
	Group       *string `json:"group" db:"group"`
	Song        *string `json:"song" db:"song"`
	ReleaseDate *string `json:"release_date" db:"release_date"`
	Text        *string `json:"text" db:"text"`
	Patronymic  *string `json:"patronymic" db:"patronymic"`
}
