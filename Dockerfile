FROM golang:1.17 as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /data/build
COPY . .
RUN go build -o app

From golang:1.17-alpine3.15 as prod
WORKDIR /data/prod
COPY --from=builder /data/build/app .
EXPOSE 8081 8082 8083

ENTRYPOINT ["./app"]