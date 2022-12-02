FROM golang:alpine

WORKDIR /go/src/hertz
COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
RUN go env
RUN go mod tidy
RUN go build -o server .

FROM alpine:latest

WORKDIR /go/src/hertz

COPY --from=0 /go/src/hertz ./

EXPOSE 8888

ENTRYPOINT ./server -c config.docker.yaml
