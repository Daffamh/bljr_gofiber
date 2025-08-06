# Gunakan image resmi Golang
FROM golang:alpine

# Buat direktori kerja di dalam container
WORKDIR /app

# Copy konfigurasi Go ke dalam container dan download dependency
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy seluruh file kode (termasuk main.go) ke dalam container
COPY . .

# Build aplikasi
RUN go build -o app .

# Jalankan aplikasi
CMD ["./app"]
