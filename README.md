# 🧱 Clean Architecture Orders System

## 📖 Descrição

Este projeto é uma implementação de um sistema de pedidos (Orders System) baseado na **Clean Architecture**, e no projeto já criado pelo professor no módulo. 

- ✅ **REST API** para criação e listagem de pedidos  
- ✅ **gRPC** para criação e listagem de pedidos  
- ✅ **GraphQL** para listagem de pedidos  

O sistema utiliza **MySQL** como banco de dados relacional e **RabbitMQ** como sistema de mensageria assíncrona.

---

## 🛠️ Tecnologias Utilizadas

- **Linguagem:** Go (Golang)  
- **Banco de Dados:** MySQL  
- **Mensageria:** RabbitMQ  
- **Orquestração:** Docker e Docker Compose  
- **Protocolos/API:** REST, gRPC e GraphQL  

---

## ▶️ Como Rodar o Projeto

### ✅ Pré-requisitos

- Docker  
- Docker Compose  

### 🚀 Subindo o ambiente

```bash
git https://github.com/fjgmelloni/20-CleanArch.git
cd 20-CleanArch

# Suba os serviços
docker-compose up --build
```

---

## 🌐 Endpoints

### 🔹 REST API

- **Criar Pedido**
  ```
  POST /order
  ```
  **Payload exemplo:**
  ```json
  {
    "id": "1",
    "price": 100.0,
    "tax": 10.0
  }
  ```

- **Listar Pedidos**
  ```
  GET /orders
  ```

---

### 🔹 gRPC

- **Criar Pedido**  
  **Método:** `CreateOrder`

- **Listar Pedidos**  
  **Método:** `ListOrders`

Use um cliente como `grpcurl` ou `Postman` para testar.

---

### 🔹 GraphQL

- **Query para listar pedidos**
  ```graphql
  query {
    listOrders {
      id
      price
      tax
      finalPrice
    }
  }
  ```

Use um cliente como **Apollo Studio**, **Insomnia** ou **GraphiQL**.

---

## 🗂️ Estrutura do Projeto

```
cmd/ordersystem        -> Arquivo principal para iniciar a aplicação
internal/usecase       -> Casos de uso (regras de negócio)
internal/infra/
  ├── database          -> Repositórios e acesso ao banco de dados
  ├── grpc              -> Implementação dos serviços gRPC
  ├── web               -> Implementação dos endpoints REST
  └── graph             -> Implementação das queries GraphQL
pkg/events             -> Eventos e dispatcher
```

---

## 🗃️ Banco de Dados

A tabela `orders` será criada automaticamente ao subir o ambiente com Docker:

```sql
CREATE TABLE orders (
  id VARCHAR(255) PRIMARY KEY,
  price DECIMAL(10, 2) NOT NULL,
  tax DECIMAL(10, 2) NOT NULL,
  final_price DECIMAL(10, 2) NOT NULL
);
```

---

## 🧪 Testando a Aplicação

### 📄 REST API (com `api.http`)

Arquivo `api.http` incluído no projeto pode ser usado em ferramentas como o VS Code para testar endpoints:

```http
# Criar Pedido
POST http://localhost:8080/order
Content-Type: application/json

{
  "id": "1",
  "price": 100.0,
  "tax": 10.0
}

###

# Listar Pedidos
GET http://localhost:8080/orders
```

### 📡 gRPC

Utilize ferramentas como:

- [`grpcurl`](https://github.com/fullstorydev/grpcurl)
- Postman (com suporte gRPC)

### 🧬 GraphQL

Ferramentas sugeridas:

- [Apollo Studio](https://studio.apollographql.com/)
- [Insomnia](https://insomnia.rest/)
- [GraphQL Playground](https://github.com/graphql/graphql-playground)

---

## 🔌 Portas dos Serviços

| Serviço     | Porta  |
|-------------|--------|
| REST API    | 8080   |
| gRPC        | 50051  |
| GraphQL     | 8081   |

---

## 👨‍💻 Autor

Desenvolvido por **[Seu Nome]**
