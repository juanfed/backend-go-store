package models

import "time"

type Product struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Brand     string    `json:"brand"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
