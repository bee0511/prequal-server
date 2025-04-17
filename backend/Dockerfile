# backend/Dockerfile
FROM golang:1.20-alpine

WORKDIR /app
COPY . .

RUN go build -o backend backend.go

EXPOSE 1233
CMD ["./backend"]
