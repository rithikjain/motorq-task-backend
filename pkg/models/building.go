package models

type Building struct {
	Name     string   `json:"name" gorm:"primaryKey"`
	Location Location `json:"location" gorm:"embedded"`
}

type Location struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}
