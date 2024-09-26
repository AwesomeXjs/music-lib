package model

type Song struct {
	Id          string `json:"id" db:"id"`
	ReleaseDate string `json:"releaseDate" db:"release_date"`
	Text        string `json:"text" db:"text"`
	Patronymic  string `json:"patronymic" db:"patronymic"`
	Group       string `json:"group" db:"group_name"`
	Song        string `json:"song" db:"song"`
}

type SongCreate struct {
	Group string `json:"group" validate:"required" example:"Muse"`
	Song  string `json:"song" validate:"required" example:"Supermassive Black Hole"`
}

type SongRequest struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Patronymic  string `json:"patronymic"`
}

type SongUpdate struct {
	Group       *string `json:"group" db:"group_name"`
	Song        *string `json:"song" db:"song"`
	ReleaseDate *string `json:"release_date" db:"release_date"`
	Text        *string `json:"text" db:"text"`
	Patronymic  *string `json:"patronymic" db:"patronymic"`
}
