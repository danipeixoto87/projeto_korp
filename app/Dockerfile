# ==========================
# Stage 1 - Build
# ==========================
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copia apenas o arquivo de dependências
COPY go.mod ./

# Copia o restante do código
COPY . .

# Organiza dependências (criará go.sum quando necessário)
RUN go mod tidy

# Compila a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o http-server ./cmd/server

# ==========================
# Stage 2 - Runtime
# ==========================
FROM alpine:3.22

RUN apk add --no-cache ca-certificates

RUN adduser -D appuser

WORKDIR /app

COPY --from=builder /app/http-server .

USER appuser

EXPOSE 8080

ENTRYPOINT ["./http-server"]
