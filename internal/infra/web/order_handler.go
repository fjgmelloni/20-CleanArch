package web

import (
    "encoding/json"
    "net/http"

    "github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type WebOrderHandler struct {
    CreateOrderUseCase *usecase.CreateOrderUseCase
    ListOrdersUseCase  *usecase.ListOrdersUseCase
}

func NewWebOrderHandler(
    createOrderUseCase *usecase.CreateOrderUseCase,
    listOrdersUseCase *usecase.ListOrdersUseCase,
) *WebOrderHandler {
    return &WebOrderHandler{
        CreateOrderUseCase: createOrderUseCase,
        ListOrdersUseCase:  listOrdersUseCase,
    }
}

func (h *WebOrderHandler) List(w http.ResponseWriter, r *http.Request) {
    output, err := h.ListOrdersUseCase.Execute()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(output)
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
    var dto usecase.OrderInputDTO
    err := json.NewDecoder(r.Body).Decode(&dto)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    output, err := h.CreateOrderUseCase.Execute(dto)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(output)
}