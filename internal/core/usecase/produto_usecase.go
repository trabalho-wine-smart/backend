package usecase

import (
	"github/arthur-psp/wine-smart/internal/core/domain"
	"github/arthur-psp/wine-smart/internal/infra/repository"
)

type ProdutoUseCase interface {
	Adiciona(p domain.Produto) error
	Lista() ([]domain.Produto, error)
}

type produtoUseCase struct {
	produtoRepo repository.ProdutoRepository
}

func (p *produtoUseCase) Adiciona(produto domain.Produto) error {
	err := p.produtoRepo.Adiciona(&produto)
	if err != nil {
		return err
	}
	return nil
}

func (p *produtoUseCase) Lista() ([]domain.Produto, error) {
	produtos, err := p.produtoRepo.ListaProdutos()
	if err != nil {
		return nil, err
	}

	return produtos, nil
}

// func (usecase *produtoUseCase) BuscaProdutoPeloID(id int64) (*domain.Produto, error) {
// 	produto, err := usecase.produtoReader.BuscaPeloID(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return produto, nil
// }

// func (usecase *produtoUseCase) AtualizaProduto(p *domain.Produto, id int64) (*domain.Produto, error) {
// 	produto, err := usecase.produtoWriter.Atualiza(p, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return produto, nil
// }

// func (usecase *produtoUseCase) RemoveProduto(id int64) error {
// 	err := usecase.produtoWriter.Deleta(id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func NewProdutoUseCase(repo repository.ProdutoRepository) ProdutoUseCase{
	return &produtoUseCase{
		produtoRepo: repo,
	}
}
