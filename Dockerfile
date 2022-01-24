#FROM golang:1.17-alpine3.15 AS builder
FROM golang:1.17 AS builder

WORKDIR /data/build
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -installsuffix cgo  -o app
RUN go build -ldflags="-s -w" -installsuffix cgo   -o jim_service cmd/main/main.go

#FROM alpine:3.15 AS final
FROM golang:1.17 AS final
WORKDIR /data/build

COPY --from=builder /data/build/config.yaml .
RUN mkdir log && touch log/rpc.log && chmod -R 777 log
RUN ls
COPY --from=builder /data/build/app .
COPY --from=builder /data/build/jim_service .

EXPOSE 9093

ENTRYPOINT ["./app"]
