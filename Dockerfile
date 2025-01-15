FROM golang:1.22.6

WORKDIR /app

COPY . .

RUN go build -o main cmd/api/main.go

EXPOSE 8080

CMD ["./main"]
