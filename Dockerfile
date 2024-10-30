FROM golang:alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /build
ADD . .
RUN go mod download
RUN go build -o main

FROM scratch
WORKDIR /app
COPY --from=builder /build/main /app

CMD ["./main"]
