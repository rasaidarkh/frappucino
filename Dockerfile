FROM golang:1.22.6

WORKDIR /app

COPY . .

RUN go build -o main cmd/myapp/main.go

EXPOSE 5433

CMD ["./main"]