package repository

import (
	"fmt"
	"github/arthur-psp/wine-smart/internal/core/domain"

	"github.com/jmoiron/sqlx"
)

type LeituraRepository interface {
	SetarNovaTemperatura (l *domain.Leituras) error
	ListaLeituras() ([]domain.Leituras, error)
}

type leituraRepository struct {
	DB *sqlx.DB
}

func NewLeituraRepository(db *sqlx.DB) LeituraRepository {
	return &leituraRepository{
		DB: db,
	}
}


func (leituraRepository *leituraRepository) SetarNovaTemperatura(l *domain.Leituras) error {
	query := `INSERT INTO leituras (temperatura, ligado, timestamp, tipo_de_vinho) VALUES ($1, $2, $3, $4) RETURNING id_temperatura`
	fmt.Println("tentando inserir: ", l.Temperatura, l.Ligado, l.Timestamp, l.TipoVinho)
	err := leituraRepository.DB.QueryRow(query, l.Temperatura, l.Ligado, l.Timestamp, l.TipoVinho).Scan(&l.ID)
	if err != nil {
		return err
	}
	return nil
}

func (leituraRepository *leituraRepository) ListaLeituras() ([]domain.Leituras, error){
	var leituras []domain.Leituras

	query, err := leituraRepository.DB.Prepare("SELECT * FROM leituras")
	if err != nil {
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var leitura domain.Leituras
		err := rows.Scan(
			&leitura.ID,
			&leitura.Temperatura,
			&leitura.Ligado,
			&leitura.Timestamp,
			&leitura.TipoVinho,
		)
		if err != nil {
			return nil, err
		}
		leituras = append(leituras, leitura)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return leituras, nil
}