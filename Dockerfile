FROM golang:1.22.6

WORKDIR /app

COPY . .

RUN go build -o main cmd/myapp/myapp.go

# Set environment variables
ENV DB_HOST=db
ENV DB_USER=latte
ENV DB_PASSWORD=latte
ENV DB_NAME=frappuccino
ENV DB_PORT=5432

EXPOSE 8080

CMD ["./main"]
