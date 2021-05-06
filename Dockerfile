FROM golang:latest AS stage1-buildphase
LABEL maintainer "Iovesophy"
WORKDIR /go/src
COPY *.go .
COPY test.txt .
COPY start.sh .
RUN go test main_test.go main.go -v
RUN go build -o tail main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=stage1-buildphase /go/src/* .
CMD ["./start.sh"]
