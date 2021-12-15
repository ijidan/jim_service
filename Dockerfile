FROM golang:1.17-alpine3.15 AS builder

WORKDIR /data/build
ENV GO111MODULE=ON \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -installsuffix cgo  -o app

FROM alpine:3.15 AS final
WORKDIR /data/prod
COPY --from=builder /data/build/config.yaml /data/build/config.yaml
COPY --from=builder /data/build/app .
EXPOSE 8081 8082 8083

ENTRYPOINT ["./app"]