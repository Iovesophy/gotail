FROM golang:latest
LABEL maintainer "Iovesophy"
WORKDIR /go/src
COPY main.go .
COPY main_test.go .
COPY test.txt .
COPY start.sh .
RUN go test main_test.go main.go -v
RUN go build main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates && \
    apk add bash
WORKDIR /root
COPY --from=0 /go/src/main .
COPY --from=0 /go/src/test.txt .
COPY --from=0 /go/src/start.sh .
CMD ["./start.sh"]
