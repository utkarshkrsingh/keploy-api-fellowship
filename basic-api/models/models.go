package models

import "gorm.io/gorm"

type WatchList struct {
	gorm.Model
	Name          string `gorm:"not null" json:"name"`
	TotalEpisodes int    `gorm:"not null" json:"totalepisodes"`
	TotalWatched  int    `gorm:"not null" json:"totalwatched"`
	Status        string `json:"status"`
	Type          string `json:"type"`
}
