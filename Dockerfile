FROM golang:1.21.6-alpine AS builder

LABEL maintainer="Muhammet Hadi KAMAT <muhammedhhadikamat@gmail.com>"

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /build

COPY go.* .

RUN go mod tidy

COPY . .

RUN go build -v -o server .

EXPOSE 3000

FROM scratch

COPY --from=builder ["/build/server", "/build/.env", "/"]

CMD ["/server"]