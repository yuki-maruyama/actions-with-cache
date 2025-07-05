package models

import "time"

type Response struct {
	Message   string    `json:"message"`
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}