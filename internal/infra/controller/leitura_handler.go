package controller

import (
	"encoding/json"
	"fmt"
	"github/arthur-psp/wine-smart/internal/core/domain"
	"github/arthur-psp/wine-smart/internal/core/usecase"
	"net/http"
)

type leituraHandler struct{
	leituraUsecase usecase.LeituraUseCase
}

func NewLeituraController(leituraUsecase usecase.LeituraUseCase) *leituraHandler {
	return &leituraHandler{
		leituraUsecase: leituraUsecase,
	}
}

func (controller *leituraHandler) Seta(response http.ResponseWriter, request *http.Request) {
	var leitura domain.Leituras
	fmt.Print("antes do decode: ", leitura)
	err := json.NewDecoder(request.Body).Decode(&leitura)

	if err != nil {
		fmt.Println("erro ao decodificar json: ", err)
		http.Error(response, "erro ao decodificar json: ", http.StatusBadRequest)
		return
	}
	fmt.Println("depois do decode: ", leitura)

	err = controller.leituraUsecase.SetarNovaTemperatura(leitura)
	if err != nil {
		http.Error(response, "erro ao adicionar produto: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(leitura); err != nil {
		http.Error(response, "erro ao decodificar JSON", http.StatusInternalServerError)
		return
	} 
}

func (controller *leituraHandler) Lista(response http.ResponseWriter, request *http.Request) {
	leituras, err := controller.leituraUsecase.ListarLeituras()
	if err != nil {
		http.Error(response, "erro ao listar leituras: ", http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(leituras)
}