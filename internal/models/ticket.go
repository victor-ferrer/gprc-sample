package models

import "time"

type Ticket struct {
	ID           int       `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	PurchaseDate time.Time `json:"purchase_date"`
	Amount       float64   `json:"amount"`
	Currency     string    `json:"currency"`
	Labels       []string  `json:"labels"`
	File         string    `json:"file"`
}
