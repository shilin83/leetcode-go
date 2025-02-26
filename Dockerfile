FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "test", "-v", "./..."]
