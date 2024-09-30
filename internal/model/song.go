package model

// Song base model for song
type Song struct {
	ID          string `json:"id" db:"id"`
	Song        string `json:"song" db:"song" example:"Supermassive Black Hole"`
	Group       string `json:"group" db:"group_name" example:"Muse"`
	ReleaseDate string `json:"releaseDate" db:"release_date" default:"NOT FOUND" example:"16.07.2006"`
	Text        string `json:"text" db:"text" default:"NOT FOUND" example:"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"`
	Link        string `json:"link" db:"link" default:"NOT FOUND" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw"`
}

// SongCreate Song create model
type SongCreate struct {
	Song  string `json:"song" validate:"required" example:"Supermassive Black Hole"`
	Group string `json:"group" validate:"required" example:"Muse"`
}

// SongRequest Song request model for song
type SongRequest struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// SongUpdate update model for song
type SongUpdate struct {
	Song        *string `json:"song" db:"song"  example:"Supermassive Black Hole"`
	Group       *string `json:"group" db:"group_name" example:"Muse"`
	ReleaseDate *string `json:"releaseDate" db:"release_date" example:"16.07.2006"`
	Text        *string `json:"text" db:"text" example:"Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"`
	Link        *string `json:"link" db:"link" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw"`
}
