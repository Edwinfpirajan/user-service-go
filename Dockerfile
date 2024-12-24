# 1. Usa una imagen base de Go
FROM golang:1.23 as builder

# 2. Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# 3. Copia los archivos del proyecto al contenedor
COPY . .

# 4. Instala las dependencias
RUN go mod tidy

# 5. Compila el binario de la aplicación
RUN go build -o user-service ./cmd/main.go

# 6. Usa una imagen base ligera para el entorno de producción
FROM debian:buster-slim

# 7. Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# 8. Copia el binario compilado desde la etapa anterior
COPY --from=builder /app/user-service .

# 9. Expone el puerto de la aplicación
EXPOSE 8080

# 10. Define el comando para ejecutar la aplicación
CMD ["./user-service"]
