package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/volli1704/prime_api/internal/number/handler"
	"github.com/volli1704/prime_api/internal/number/usecase"
	"github.com/volli1704/prime_api/pkg/primechecker"
)

func main() {
	r := mux.NewRouter()

	checker := primechecker.PrimeChecker{}
	numberUseCase := usecase.NewNumberUseCase(checker)
	handler.NewNumberHandler(r, numberUseCase)

	http.ListenAndServe(":8080", r)
}
