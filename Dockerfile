# 1. Use the official Golang image as the base
FROM golang:1.22 AS builder

# 2. Set the working directory
WORKDIR /app

# 3. Copy go.mod first and download deps
COPY go.mod ./
RUN go mod download

# 4. Copy the rest of the source code
COPY . .

# 5. Build the Go application
RUN go build -o ascii-art-web-dockerize ./main.go

# 6. Use a smaller base image for the final image
FROM debian:bookworm-slim

# 7. Set working directory
WORKDIR /app

# 8. Copy the binary from builder
COPY --from=builder /app/ascii-art-web-dockerize .

# 9. Copy required folders (your templates, static, etc.)
COPY --from=builder /app/internal /app/internal

# 10. Expose port 8080
EXPOSE 8080

# 11. Metadata
LABEL maintainer="basel29ali@outlook.com"
LABEL project="ascii-art-web-dockerize"
LABEL version="1.0"

# 12. Run the binary
CMD ["./ascii-art-web-dockerize"]
