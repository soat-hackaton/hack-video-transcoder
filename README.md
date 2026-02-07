# hack-video-transcoder

# Executar projeto

go run cmd/api/main.go

# Atualizar swagger

swag init -g cmd/api/main.go

# Atualizar modulos

go mod tidy

# Documentação

http://0.0.0.0:8080/docs/index.html

# Tests

go test ./... -cover
