# -------- Stage 1: Build --------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Dependências necessárias para build
RUN apk add --no-cache git

# Copiar go.mod primeiro (cache de dependências)
COPY go.mod go.sum ./
RUN go mod download

# Copiar restante do código
COPY . .

# Build do binário
RUN go build -o api cmd/api/main.go

# -------- Stage 2: Runtime --------
FROM alpine:3.20

WORKDIR /app

# Instalar ffmpeg (necessário em runtime)
RUN apk add --no-cache ffmpeg

# Copiar binário apenas (imagem MUITO menor)
COPY --from=builder /app/api .

# Criar diretórios necessários
RUN mkdir -p uploads outputs temp

EXPOSE 8080

CMD ["./api"]
