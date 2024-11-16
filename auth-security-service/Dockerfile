# Base image
FROM golang:1.17-alpine

# Çalışma dizini
WORKDIR /app

# Bağımlılıkları kopyala ve yükle
COPY go.mod go.sum ./
RUN go mod download

# Uygulama dosyalarını kopyala
COPY . .

# Uygulamayı derle
RUN go build -o main cmd/main.go

# Portları aç
EXPOSE 8443

# Komutu çalıştır
CMD ["./main"]
