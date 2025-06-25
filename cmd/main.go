package main

import (
	"github/arthur-psp/wine-smart/internal/core/usecase"
	"github/arthur-psp/wine-smart/internal/infra/controller"
	"github/arthur-psp/wine-smart/internal/infra/repository"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, request *http.Request) {
		enableCors(resp)

		if request.Method == "OPTIONS" {
			return
		}

		next(resp, request)
	}
}

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=wine_smart sslmode=disable password=postgrespassword host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("successfully connected.")
	}

	// produto
	produtoRepo := repository.NewProdutoRepository(db)
	produtoUseCase := usecase.NewProdutoUseCase(produtoRepo)
	produtoController := controller.NewProdutoController(produtoUseCase)

	http.HandleFunc("/produtos", produtoController.Processa)
	http.HandleFunc("/listaprodutos", produtoController.Lista)

	//leituras
	leituraRepo := repository.NewLeituraRepository(db)
	leituraUseCase := usecase.NewLeituraUseCase(leituraRepo)
	leituraController := controller.NewLeituraController(leituraUseCase)

	http.HandleFunc("/leitura", withCORS(leituraController.Seta))
	http.HandleFunc("/leituras", withCORS(leituraController.Lista))

	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
