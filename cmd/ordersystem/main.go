package main

import (
    "database/sql"
    "fmt"
    "net"
    "net/http"

    graphql_handler "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/devfullcycle/20-CleanArch/configs"
    "github.com/devfullcycle/20-CleanArch/internal/event/handler"
    "github.com/devfullcycle/20-CleanArch/internal/infra/database"
    "github.com/devfullcycle/20-CleanArch/internal/infra/graph"
    "github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
    "github.com/devfullcycle/20-CleanArch/internal/infra/grpc/service"
    "github.com/devfullcycle/20-CleanArch/internal/infra/web/webserver"
    "github.com/devfullcycle/20-CleanArch/internal/usecase"
    eventpkg "github.com/devfullcycle/20-CleanArch/internal/event" // Renamed to avoid conflict
    "github.com/devfullcycle/20-CleanArch/pkg/events"
    "github.com/streadway/amqp"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    // mysql
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Load configurations
    configs, err := configs.LoadConfig(".")
    if err != nil {
        panic(err)
    }

    // Database connection
    db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // RabbitMQ setup
    rabbitMQConn, rabbitMQChannel := getRabbitMQChannel()
    defer rabbitMQConn.Close()  // Ensure RabbitMQ connection is closed
    defer rabbitMQChannel.Close() // Ensure RabbitMQ channel is closed

    // Event dispatcher setup
    eventDispatcher := events.NewEventDispatcher()
    eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
        RabbitMQChannel: rabbitMQChannel,
    })

    // Instantiate repository
    orderRepository := database.NewOrderRepository(db)

    // Instantiate use cases
    orderCreatedEvent := eventpkg.NewOrderCreated()
    createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, orderCreatedEvent, eventDispatcher)
    listOrdersUseCase := usecase.NewListOrdersUseCase(orderRepository)

    // Web server setup
    webserverInstance := webserver.NewWebServer(configs.WebServerPort)
    webOrderHandler := webserver.NewWebOrderHandler(createOrderUseCase, listOrdersUseCase)
    webserverInstance.AddHandler("/order", webOrderHandler.Create, "POST")
    webserverInstance.AddHandler("/order", webOrderHandler.List, "GET") // Add ListOrders route
    fmt.Println("Starting web server on port", configs.WebServerPort)
    go webserverInstance.Start()

    // gRPC server setup
    grpcServer := grpc.NewServer()
    orderService := service.NewOrderService(createOrderUseCase, listOrdersUseCase)
    pb.RegisterOrderServiceServer(grpcServer, orderService)
    reflection.Register(grpcServer)

    fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
    lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
    if err != nil {
        panic(err)
    }
    go grpcServer.Serve(lis)

    // GraphQL server setup
    srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
        CreateOrderUseCase: createOrderUseCase,
        ListOrdersUseCase:  listOrdersUseCase,
    }}))
    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
    http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() (*amqp.Connection, *amqp.Channel) {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        panic(err)
    }
    ch, err := conn.Channel()
    if err != nil {
        conn.Close() // Close the connection if channel creation fails
        panic(err)
    }
    return conn, ch
}