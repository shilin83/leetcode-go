FROM golang:1.23.4-alpine3.21

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "test", "-v", "./tests/..."]
