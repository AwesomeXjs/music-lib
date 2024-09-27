package model

type Song struct {
	Id          string `json:"id" db:"id"`
	ReleaseDate string `json:"releaseDate" db:"release_date" default:"NOT FOUND" example:"16.07.2006"`
	Text        string `json:"text" db:"text" default:"NOT FOUND" example:"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"`
	Patronymic  string `json:"patronymic" db:"patronymic" default:"NOT FOUND" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw"`
	Group       string `json:"group" db:"group_name" example:"Muse"`
	Song        string `json:"song" db:"song" example:"Supermassive Black Hole"`
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
	Group       *string `json:"group" db:"group_name" example:"Muse"`
	Song        *string `json:"song" db:"song"  example:"Supermassive Black Hole"`
	ReleaseDate *string `json:"release_date" db:"release_date" example:"16.07.2006"`
	Text        *string `json:"text" db:"text" example:"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"`
	Patronymic  *string `json:"patronymic" db:"patronymic" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw"`
}
