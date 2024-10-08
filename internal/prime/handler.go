package prime

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"prime/types"
)

type PrimeHandler interface {
	SimulationPrice(res http.ResponseWriter, req *http.Request)
	AddItemPrime(res http.ResponseWriter, req *http.Request)
	RemoveItemPrime(res http.ResponseWriter, req *http.Request)
}

type primeHandler struct {
	service PrimeService
}

func NewPrimeHandler(service PrimeService) PrimeHandler {
	return &primeHandler{service: service}
}

func (h *primeHandler) SimulationPrice(res http.ResponseWriter, req *http.Request) {
	var input types.PriceSimulationInput
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		slog.Error(err.Error())
		writeError(http.StatusBadRequest, "Invalid input", res)
		return
	}

	result, err := h.service.SimulationPrice(req.Context(), input)
	if err != nil {
		slog.Error(err.Error())
		writeError(http.StatusInternalServerError, "Failed to simulate price", res)
		return
	}

	writeJSON(res, result)
}

func (h *primeHandler) AddItemPrime(res http.ResponseWriter, req *http.Request) {
	// Implementation of AddItemPrime
}

func (h *primeHandler) RemoveItemPrime(res http.ResponseWriter, req *http.Request) {
	// Implementation of RemoveItemPrime
}

func writeError(status int, message string, res http.ResponseWriter) {
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(map[string]string{"error": message})
}

func writeJSON(res http.ResponseWriter, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(data); err != nil {
		writeError(http.StatusInternalServerError, "Failed to write response", res)
	}
}
