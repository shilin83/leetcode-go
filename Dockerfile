FROM golang:1.23.5-alpine

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "test", "-v", "./..."]
