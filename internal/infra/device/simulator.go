package device

import (
	"encoding/json"
	"fmt"
	"github/arthur-psp/wine-smart/internal/core/domain"
	"github/arthur-psp/wine-smart/internal/core/usecase"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type simulador struct {
	leituraUsecase usecase.LeituraUseCase
	temperaturaAlvo int
	temperaturaAtual int
}

func NewSimulador(leituraUsecase usecase.LeituraUseCase) *simulador{
	rand.Seed(time.Now().UnixNano())
	inicial := 6 + rand.Intn(11)
	return &simulador{
		leituraUsecase: leituraUsecase,
		temperaturaAtual: inicial,
		temperaturaAlvo: inicial,
	}
}

func gerarTemperaturaFake() int {
	min := 6
	max := 16
	temperaturaAtual := min + rand.Intn(max - min)
	return temperaturaAtual 
}

func (controller *simulador) Seta(response http.ResponseWriter, request *http.Request) {
	temperatura := gerarTemperaturaFake()

	leitura := domain.Leituras{
		Temperatura: temperatura,
		Ligado: true,
		Timestamp: time.Now(),
		TipoVinho: "Vinho Tinto",
	}
	fmt.Println("tmepe gerada: ", temperatura)

	err := controller.leituraUsecase.SetarNovaTemperatura(leitura)
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

func (controller *simulador) Lista(response http.ResponseWriter, request *http.Request) {
	leituras, err := controller.leituraUsecase.ListarLeituras()
	if err != nil {
		http.Error(response, "erro ao listar leituras: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(leituras)
}

func (s *simulador) RegularTemperatura(resp http.ResponseWriter, request *http.Request) {
	var dados struct {
		Temperatura int `json:"temperatura"`

	}

	err := json.NewDecoder(request.Body).Decode(&dados)
	if err != nil {
		http.Error(resp, "erro ao decodificar json: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if dados.Temperatura < 6 || dados.Temperatura > 16 {
		http.Error(resp, "temperatura fora no intervalo permitido. ", http.StatusBadRequest)
		return
	}

	s.temperaturaAlvo = dados.Temperatura
	fmt.Println("temperatura alvo: ", s.temperaturaAlvo)

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("Temperatura alvo definida com sucesso!"))
}

func (s *simulador) Start() {
	go func() {
		for {
			if s.temperaturaAtual < s.temperaturaAlvo {
				s.temperaturaAtual++
			} else if s.temperaturaAtual > s.temperaturaAlvo {
				s.temperaturaAtual--
			} else {
				// Já chegou na temperatura alvo
				time.Sleep(7 * time.Second)
				continue
			}

			// Salva no banco
			leitura := domain.Leituras{
				Temperatura: s.temperaturaAtual,
				Ligado:      true,
				Timestamp:   time.Now(),
				TipoVinho: "Vinho tinto",
			}

			fmt.Println("Atualizando temperatura:", s.temperaturaAtual)

			if err := s.leituraUsecase.SetarNovaTemperatura(leitura); err != nil {
				log.Println("Erro ao salvar leitura:", err)
			}

			time.Sleep(5 * time.Second)
		}
	}()
}
