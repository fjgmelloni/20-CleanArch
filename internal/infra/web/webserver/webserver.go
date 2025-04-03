package webserver

import (
    "encoding/json"
    "net/http"

    "github.com/devfullcycle/20-CleanArch/internal/usecase"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
    Router        chi.Router
    WebServerPort string
}

func NewWebServer(port string) *WebServer {
    return &WebServer{
        Router:        chi.NewRouter(),
        WebServerPort: port,
    }
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc, method string) {
    s.Router.Method(method, path, handler)
}

func (s *WebServer) Start() {
    s.Router.Use(middleware.Logger)
    http.ListenAndServe(":"+s.WebServerPort, s.Router)
}

type WebOrderHandler struct {
    CreateOrderUseCase *usecase.CreateOrderUseCase
    ListOrdersUseCase  *usecase.ListOrdersUseCase
}

func NewWebOrderHandler(createOrderUseCase *usecase.CreateOrderUseCase, listOrdersUseCase *usecase.ListOrdersUseCase) *WebOrderHandler {
    return &WebOrderHandler{
        CreateOrderUseCase: createOrderUseCase,
        ListOrdersUseCase:  listOrdersUseCase,
    }
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
    var input usecase.OrderInputDTO
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    output, err := h.CreateOrderUseCase.Execute(input)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(output)
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