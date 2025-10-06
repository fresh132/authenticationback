FROM golang:1.25.1 AS backend

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

EXPOSE 9091

CMD ["./main"]
