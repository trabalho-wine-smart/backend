package controller

import (
	"encoding/json"
	"fmt"
	"github/arthur-psp/wine-smart/internal/core/domain"
	"github/arthur-psp/wine-smart/internal/core/usecase"
	"net/http"
)

type produtoController struct {
	produtoUseCase usecase.ProdutoUseCase
}

func NewProdutoController(produtoUsecase usecase.ProdutoUseCase) *produtoController {
	return &produtoController{
		produtoUseCase: produtoUsecase,
	}
}

func (controller *produtoController) Processa(response http.ResponseWriter, request *http.Request) {
	var produto domain.Produto
	fmt.Println("antes do decode ", produto)
	err := json.NewDecoder(request.Body).Decode(&produto)
	
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		http.Error(response, "Erro ao decodificar JSON: ", http.StatusBadRequest)
		return
	}

	fmt.Println("depois do decode ", produto)


	err = controller.produtoUseCase.Adiciona(produto)
	if err != nil {
		http.Error(response, "Erro ao adicionar produto: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(produto); err != nil {
		http.Error(response, "Erro ao codificar JSON", http.StatusInternalServerError)
		return
	}
	
}



func (controller *produtoController) Lista(res http.ResponseWriter, request *http.Request) {
	produtos, error := controller.produtoUseCase.Lista()
	if error != nil {
		http.Error(res, "Erro ao listar produtos", http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(produtos)

	fmt.Println("code ", produtos)
}

// func (controller *produtoController) BuscaPeloID(id int64, res http.ResponseWriter, request *http.Request) {
// 	produto, err := controller.produtoUseCase.BuscaProdutoPeloID(id)
// 	if err != nil {
// 		http.Error(res, "Erro ao buscar produto pelo ID", http.StatusInternalServerError)
// 		return
// 	}
// 	res.WriteHeader(http.StatusOK)
// 	json.NewEncoder(res).Encode(produto)
// 	fmt.Println("produto pelo id encode: ", produto)
// }

// func (controller *produtoController) AtualizaProduto(res http.ResponseWriter, request *http.Request) {
// 	var produto domain.Produto

// 	// Decode o JSON do body para a struct
// 	if err := json.NewDecoder(request.Body).Decode(&produto); err != nil {
// 		http.Error(res, "Erro ao decodificar JSON", http.StatusBadRequest)
// 		return
// 	}

// 	// O ID pode estar dentro do próprio JSON
// 	id := produto.ID
// 	if id == 0 {
// 		http.Error(res, "ID do produto não fornecido", http.StatusBadRequest)
// 		return
// 	}

// 	// Chama o usecase
// 	produtoAtualizado, err := controller.produtoUseCase.AtualizaProduto(&produto, id)
// 	if err != nil {
// 		http.Error(res, "Erro ao atualizar produto", http.StatusInternalServerError)
// 		return
// 	}

// 	// Retorna o produto atualizado como JSON
// 	res.WriteHeader(http.StatusOK)
// 	json.NewEncoder(res).Encode(produtoAtualizado)
// 	fmt.Println("Produto atualizado:", produtoAtualizado)

// 	// produto, err := controller.produtoUseCase.AtualizaProduto(p, id)
// 	// if err != nil {
// 	// 	http.Error(res, "Erro ao atualizar produto: ", http.StatusInternalServerError)
// 	// 	return
// 	// }

// 	// //res.WriteHeader(http.StatusOK)
// 	// json.NewEncoder(res).Encode(produto)
// 	// fmt.Println("Produto encodade: ", produto)
// }

// func (controller *produtoController) RemoveProduto(id int64, res http.ResponseWriter, request http.Request) {
// 	err := controller.produtoUseCase.RemoveProduto(id)
// 	if err != nil {
// 		http.Error(res, "Erro ao remover produto: ", http.StatusInternalServerError)
// 		return
// 	}
// 	res.WriteHeader(http.StatusOK)
// 	json.NewEncoder(res).Encode(err)
// 	fmt.Println("Produto encodade: ", err)
// }


