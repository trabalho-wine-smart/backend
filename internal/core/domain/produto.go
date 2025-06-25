package domain

type Produto struct {
	ID        int64   `json:"id_produto"`
	Nome      string  `json:"nome"`
	Preco     float64 `json:"preco"`
	Descricao string  `json:"descricao"`
	Status    string  `json:"status"`
}

// type ProdutoWriter interface {
// 	Adiciona(p Produto) error
// 	// Atualiza(p *Produto, id int64) (*Produto, error)
// 	// Deleta(id int64) error
// }

// type ProdutoReader interface {
// 	ListaProdutos() ([]Produto, error)
// 	BuscaPeloID(id int64) (*Produto, error)
// }
