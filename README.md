# 🚀 Encurtador de URL

Este projeto é uma simples API encurtadora de URL construída usando a linguagem de programação Go e o banco de dados Redis para armazenamento de dados. Ela permite que você encurte URLs longas em URLs curtas.

## Tecnologias Utilizadas

- **Go**: A aplicação é desenvolvida em Go, uma linguagem de programação compilada e eficiente.
- **Gin**: O framework web Gin é utilizado para facilitar o desenvolvimento do servidor HTTP.
- **Redis**: O Redis é utilizado como um banco de dados em memória para armazenar as informações de URL encurtadas.

## Como Usar

1. Clone o repositório:

    ```bash
    git clone https://github.com/christoffer1009/url-shortener-go.git
    ```

2. Entre no diretório do projeto:

    ```bash
    cd url-shortener-go
    ```

3. Configure as variáveis de ambiente no arquivo `.env`:

    ```plaintext
    REDIS_HOST=seu-host-redis
    REDIS_PORT=6379
    REDIS_PASSWORD=sua_senha_redis
    APP_PORT=8080
    ```

4. Execute a aplicação:

    ```bash
    go run main.go
    ```

5. Acesse a aplicação em [http://localhost:8080](http://localhost:8080).

## API Endpoints

- **Encurtar URL:**
  - Endpoint: `POST /encurtar`
  - Corpo da Requisição: `{ "original_url": "URL_LONGA_A_SER_ENCURTADA" }`

- **Redirecionar URL:**
  - Endpoint: `GET /{código}`
  - Redireciona para a URL original associada ao código fornecido.

## Contribuição

Sinta-se à vontade para abrir issues e enviar pull requests para melhorar este projeto.
