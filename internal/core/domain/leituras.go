package domain

import "time"

type Leituras struct {
	ID	int64	`json:"id_temperatura"`
	Temperatura float64 `json:"temperatura"`
	Ligado bool `json:"ligado"`
	Timestamp time.Time `json:"timestamp"`
}