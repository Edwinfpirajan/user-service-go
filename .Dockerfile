# Etapa 1: Construcción
FROM golang:1.23 AS builder

# Configurar el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos de dependencias y descargar los módulos
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente al contenedor
COPY . .

# Construir el ejecutable
RUN go build -o main ./cmd/main.go

# Etapa 2: Imagen final (más ligera)
FROM debian:bullseye-slim

# Instalar librerías necesarias para ejecutar el binario
RUN apt-get update && apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Crear un directorio para la aplicación
WORKDIR /app

# Copiar el ejecutable desde la etapa de construcción
COPY --from=builder /app/main .

# Copiar el archivo .env al contenedor
COPY .env .

# Exponer el puerto que utiliza la aplicación (ajústalo según sea necesario)
EXPOSE 8000

# Comando para ejecutar la aplicación
CMD ["./main"]
