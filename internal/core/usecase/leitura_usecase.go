package usecase

import (
	"github/arthur-psp/wine-smart/internal/core/domain"
	"github/arthur-psp/wine-smart/internal/infra/repository"
)

type LeituraUseCase interface{
	SetarNovaTemperatura (l domain.Leituras) error
	ListarLeituras() ([]domain.Leituras, error)
}

type leituraUseCase struct {
	leituraRepo repository.LeituraRepository
}

func NewLeituraUseCase(repository repository.LeituraRepository) LeituraUseCase {
	return &leituraUseCase{
		leituraRepo: repository,
	}
}

func (leituraUsecase *leituraUseCase) SetarNovaTemperatura(l domain.Leituras) error {
	err := leituraUsecase.leituraRepo.SetarNovaTemperatura(&l)
	if err != nil {
		return err
	}
	return nil
}

func (leituraUsecase *leituraUseCase) ListarLeituras() ([]domain.Leituras, error) {
	leituras, err := leituraUsecase.leituraRepo.ListaLeituras()
	if err != nil {
		return nil, err
	}

	return leituras, nil
}