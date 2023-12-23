# Imagem base do Golang
FROM golang:alpine

# Configurar o diretório de trabalho
WORKDIR /app

# Adicionar uma linha para criar um volume para dados da aplicação
VOLUME ["/app/data"]

# Copiar arquivos necessários
COPY . .

# Baixar as dependências do Go
RUN go get -d -v ./...

# Construir o aplicativo
RUN go build -o url-shortener .

EXPOSE 8080

# Comando para iniciar o aplicativo
CMD ["./url-shortener"]

