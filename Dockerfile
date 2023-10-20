FROM golang:1.21.3

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

CMD ["/app/main"]