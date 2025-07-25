Desafio Clean Architecture: Listagem de Orders

Clone o Repositório:
    git clone https://github.com/hicarogodinho/desafioCleanArch.git
    cd desafioCleanArch

Inicie a Aplicação e o Banco de Dados com Docker Compose:
    docker compose up --build -d


Portas e Endpoints da Aplicação:

Serviço REST (HTTP)
    Porta: 8000
    URL Base: http://localhost:8000
    
    CreateOrder:
    Método: POST
    Endpoint: /order
    Exemplo de Body RAW:
        JSON
        {
          "id":"prod_1",
          "price": 5.00,
          "tax": 0.50
        }
    
    ListOrders:
        Método: GET
        Endpoint: /orders
        Exemplo de Resposta GET:
            JSON
            [
              {
                "id":"prod_1",
                "price":4.2,
                "tax":0.42,
                "final_price":4.62
              },
              {
                "id":"prod_2",
                "price":5,
                "tax":0.5,
                "final_price":5.5
              }
            ]

Serviço gRPC
    Porta: 50051
    Serviço: OrderService
    Métodos:
        CreateOrder (para criar uma nova ordem)
        ListOrders (para listar as ordens existentes)

Serviço GraphQL
    Porta: 8080
    Endpoint do Playground/IDE: / (Acessível em http://localhost:8080)
    Endpoint da Query/Mutation: /query

    Mutation CreateOrder:
    mutation createOrder {
      createOrder(input: {id: "item_1", Price: 20.50, Tax:2.0 }) {
        id
        Price
        Tax
        FinalPrice # Você pode adicionar FinalPrice aqui se ele for retornado após a criação
      }
    }

    Query ListOrders:
    query listOrders {
      orders {
        id
        Price
        Tax
        FinalPrice
      }
    }

Arquivos de Requisições HTTP (.http)
    api/create_order.http
    api/list_orders.http