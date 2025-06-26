package domain

import "time"

type Leituras struct {
	ID	int64	`json:"id_temperatura"`
	Temperatura int `json:"temperatura"`
	Ligado bool `json:"ligado"`
	Timestamp time.Time `json:"timestamp"`
	TipoVinho string `json:"tipo_de_vinho"`
}