package usecase

import (
    "github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListOrdersOutputDTO struct {
    Orders []OrderDTO
}

type OrderDTO struct {
    ID         string  `json:"id"`
    Price      float64 `json:"price"`
    Tax        float64 `json:"tax"`
    FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
    OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
    return &ListOrdersUseCase{
        OrderRepository: orderRepository,
    }
}

func (u *ListOrdersUseCase) Execute() (*ListOrdersOutputDTO, error) {
    orders, err := u.OrderRepository.List()
    if err != nil {
        return nil, err
    }

    var dtos []OrderDTO
    for _, order := range orders {
        dtos = append(dtos, OrderDTO{
            ID:         order.ID,
            Price:      order.Price,
            Tax:        order.Tax,
            FinalPrice: order.FinalPrice,
        })
    }

    return &ListOrdersOutputDTO{Orders: dtos}, nil
}