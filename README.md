# Hack Video Transcoder

## O que é o projeto
O **hack-video-transcoder** é um microsserviço (API REST) desenvolvido em Go responsável por receber uploads de arquivos de vídeo e realizar seu processamento e transcodificação. Ele conta com uma integração ao FFmpeg para realizar as operações de processamento de mídia e opera como um dos nós no pipeline assíncrono de manipulação de vídeos.

## Principais tecnologias utilizadas
- **Go (Golang) 1.25+**: Linguagem de programação principal.
- **Gin Web Framework**: Roteador HTTP rápido para a construção da API.
- **FFmpeg**: Motor core de processamento e transcodificação de vídeos.
- **Swagger (swaggo)**: Utilizado para prover a documentação interativa da API.
- **Docker**: Containerização com builds eficientes e isolados.
- **Kubernetes**: Orquestração e manutenção do deployment do serviço (Amazon EKS).

## Pré-requisitos / Variáveis de Ambiente
### Para a Execução Local
- **Go 1.25** ou superior.
- **FFmpeg** instalado e acessível no `$PATH` do sistema.

### Variáveis para o Deploy (CI/CD)
O projeto recebe a configuração a partir de *Secrets* do GitHub para gerenciar a infraestrutura na AWS durante o CI/CD:
- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_SESSION_TOKEN`
- `AWS_REGION`
- `ECR_REPOSITORY`

Para executar o projeto localmente:
```bash
# Executar projeto
go run cmd/api/main.go

# Atualizar swagger
swag init -g cmd/api/main.go

# Atualizar módulos
go mod tidy
```
> **Documentação (Swagger)**: Acessível via `http://0.0.0.0:8080/docs/index.html` após rodar o projeto.

## Testes e Qualidade
O projeto prioriza a qualidade do código com a execução de **Testes Unitários** validados com métricas de cobertura de código.

```bash
# Executar os testes localmente
go test ./... -cover
```

## Pipeline CI/CD
Foi configurado um pipeline robusto com **GitHub Actions** (`.github/workflows/ci.yml`) focado em entrega contínua. Ele engatilha a partir de *pushes* ou *pull requests* na branch `main` e é dividido em três steps:
1. **DockerBuild**: Constrói a imagem Docker isolada (`target: builder`) e envia para o GitHub Container Registry (GHCR).
2. **UnitTest**: Puxa a imagem recém-construída e roda nativamente no container o comando de coberturas e testes `go test ./... -cover`.
3. **Deploy**: Após as validações com sucesso, constrói a imagem final e a envia para o Amazon ECR. Atualiza dinamicamente as manifestações do Kubernetes (`envsubst` em namespaces, deployments e services) instalando tudo no cluster da AWS EKS.

## Melhores práticas implementadas
- **Clean Architecture & SOLID**: O código está fortemente modularizado (`domain`, `application`, `adapters`, `infra`), separando as responsabilidades de roteamento HTTP, da regra de negócio e adaptações externas (como FFmpeg).
- **Tratamento de Logs Avançado**: Logs nativamente estruturados com **Context Propagation** utilizando `Correlation IDs` em todos os steps geradores de tarefas, facilitando rastreamento de problemas.
- **Documentação como Código**: As rotas da API estão descritas junto aos *Handlers* usando anotações Swagger auto-geráveis.
- **Infraestrutura como Código (IaC)**: Manifestos de Kubernetes mantidos e versionados no diretório `infra/` separados por responsabilidade e orquestados pelo CD.
