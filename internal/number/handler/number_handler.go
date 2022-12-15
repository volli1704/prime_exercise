package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/volli1704/prime_api/internal/number/usecase"
	"github.com/volli1704/prime_api/internal/utils"
)

// ResponseError is a basic struct of json error response
type ResponseError struct {
	Message string `json:"error"`
}

// NumberHandler handles requests to number API
type NumberHandler struct {
	useCase usecase.NumberUseCase
}

func NewNumberHandler(r *mux.Router, useCase usecase.NumberUseCase) {
	handler := NumberHandler{useCase}

	r.HandleFunc("/", handler.CheckPrimes).Methods(http.MethodPost)
}

// CheckPrimes is a handler for requests
//
// It requires json number array in POST request body to return
// json bool array with result of prime check for each number
func (n *NumberHandler) CheckPrimes(w http.ResponseWriter, r *http.Request) {
	// response should be json
	w.Header().Add("Content-Type", "application/json")

	arrBody, err := utils.JsonReaderToArray(r.Body)
	if err != nil {
		writeErr(w, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	res, err := n.useCase.FindPrimesForArray(arrBody)
	if err != nil {
		writeErr(w, err)
		return
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		writeErr(w, err)
		return
	}
	w.Write(jsonRes)
}

func writeErr(w http.ResponseWriter, err error) {
	json.NewEncoder(w).Encode(ResponseError{err.Error()})
}
