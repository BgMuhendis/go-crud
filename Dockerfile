FROM golang:1.21.6

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -v -o server

EXPOSE 3000

CMD ["./server"]