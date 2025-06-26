package models

type Record struct {
	ID              int    `json:"id" db:"id"`
	Title           string `json:"title" db:"title"`
	TotalEpisodes   int    `json:"total_episodes" db:"total_episodes"`
	WatchedEpisodes int    `json:"watched_episodes" db:"watched_episodes"`
	Type            string `json:"type" db:"type"`
	Status          string `json:"status" db:"status"`
}
