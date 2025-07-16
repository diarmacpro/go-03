# Gunakan image dasar golang
FROM golang:1.21

# Set direktori kerja
WORKDIR /app

# Salin semua file ke dalam container
COPY . .

# Unduh dependensi
RUN go mod tidy

# Build binary
RUN go build -o main .

# Jalankan binary
CMD ["./main"]
