# Etapa de compilación
FROM golang:1.23.4 as builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o user-service ./cmd/main.go

# Etapa de producción con Debian reciente
FROM debian:bookworm-slim  # Imagen base con GLIBC >= 2.34

WORKDIR /app

COPY --from=builder /app/user-service .

EXPOSE 8080

CMD ["./user-service"]
