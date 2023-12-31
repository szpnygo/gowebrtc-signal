FROM golang:1.21-alpine AS builder
WORKDIR /app

ENV GOOS linux
RUN go env -w GO111MODULE=on

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o app main.go

FROM alpine

WORKDIR /app
COPY --from=builder /app/app /app/app
COPY --from=builder /app/config.yaml /app/config.yaml

CMD ["./app"]
