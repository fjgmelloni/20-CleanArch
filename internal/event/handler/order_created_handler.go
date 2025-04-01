package handler

import (
    "encoding/json"
    "fmt"
    "sync"

    "github.com/devfullcycle/20-CleanArch/pkg/events"
    "github.com/streadway/amqp"
)

type OrderCreatedHandler struct {
    RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
    return &OrderCreatedHandler{
        RabbitMQChannel: rabbitMQChannel,
    }
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
    defer wg.Done()
    
    fmt.Printf("Order created: %v\n", event.GetPayload())
    jsonOutput, err := json.Marshal(event.GetPayload())
    if err != nil {
        fmt.Printf("Error marshaling event payload: %v\n", err)
        return
    }

    msgRabbitmq := amqp.Publishing{
        ContentType: "application/json",
        Body:        jsonOutput,
    }

    err = h.RabbitMQChannel.Publish(
        "orders",    // exchange
        "order.created", // routing key
        false,      // mandatory
        false,      // immediate
        msgRabbitmq, // message to publish
    )
    if err != nil {
        fmt.Printf("Error publishing message: %v\n", err)
    }
}