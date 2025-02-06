FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main main.go

EXPOSE 3030

CMD ["./main"]
