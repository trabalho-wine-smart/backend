package repository

import (
	"fmt"

	"github/arthur-psp/wine-smart/internal/core/domain"

	"github.com/jmoiron/sqlx"
)

type ProdutoRepository interface {
	Adiciona (p *domain.Produto) error
	ListaProdutos() ([]domain.Produto, error)
}

type produtoRepository struct {
	DB *sqlx.DB
}

func NewProdutoRepository(db *sqlx.DB) ProdutoRepository {
	return &produtoRepository{
		DB: db,	
	}
}



func (produtoRepository *produtoRepository) Adiciona(p *domain.Produto) error {
	query := `INSERT INTO produtos (nome, preco, descricao, status) VALUES ($1, $2, $3, $4) RETURNING id_produto`

	fmt.Println("Tentando inserir:", p.Nome, p.Preco, p.Status, p.Descricao)

	err := produtoRepository.DB.QueryRow(query, p.Nome, p.Preco, p.Descricao, p.Status).Scan(&p.ID)
	if err != nil {
		return err
	}
	return nil

}

func (p *produtoRepository) ListaProdutos() ([]domain.Produto, error) {
	var produtos []domain.Produto

	stmt, err := p.DB.Prepare("SELECT * FROM produtos")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var produto domain.Produto
		err := rows.Scan(
			&produto.ID,
			&produto.Nome,
			&produto.Preco,
			&produto.Status,
			&produto.Descricao,
		)
		if err != nil {
			return nil, err
		}
		produtos = append(produtos, produto)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return produtos, nil
}

// func (produtoRepository *ProdutoRepository) BuscaPeloID(id int64) (produto *domain.Produto, err error) {
// 	query := `SELECT id_produto, nome, preco, descricao, status FROM produtos WHERE id_produto = $1`
// 	row := produtoRepository.DB.QueryRow(query, id)

// 	var p domain.Produto
// 	err = row.Scan(&p.ID, &p.Nome, &p.Preco, &p.Status, &p.Descricao)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Print("jk")
// 	return &p, nil
// }

// func (produtoRepository *ProdutoRepository) Atualiza(produto *domain.Produto, id int64) (*domain.Produto, error) {
// 	query := `UPDATE produtos SET nome = $1, preco = $2, status = $3, descricao = $4 WHERE id_produto = $5 RETURNING id_produto`
// 	err := produtoRepository.DB.QueryRow(query, produto.Nome, produto.Preco, produto.Status, produto.Descricao, id).Scan(&produto.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return produto, nil
// }

// // func (produtoRepository *ProdutoRepository) Atualiza(produto *domain.Produto, id int64) (*domain.Produto, error) {
// // 	query := `UPDATE produtos SET nome = $1, preco = $2, status = $3, descricao = $4 WHERE id_produto = $5 RETURNING id_produto`
// // 	err := produtoRepository.DB.QueryRow(query, produto.Nome, produto.Preco, produto.Status, produto.Descricao, id).Scan(&produto.ID)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return produto, nil
// // }


// func (produtoRepository *ProdutoRepository) Deleta(id int64) error {
// 	query := `DELETE FROM produtos WHERE id_produto = $1`

// 	result, err := produtoRepository.DB.Exec(query, id)
// 	if err != nil {
// 		return err
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if rowsAffected == 0 {
// 		return fmt.Errorf("nenhum produto com id %d encontrado", id)
// 	}

// 	return nil
// }
