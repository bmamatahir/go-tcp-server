FROM golang:1.20-alpine

WORKDIR /app

COPY server.go .

RUN go build -o server server.go

CMD ["./server"]
