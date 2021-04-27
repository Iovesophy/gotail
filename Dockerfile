FROM golang:latest
LABEL maintainer "Iovesophy"
WORKDIR /go/src
COPY main.go .
RUN go build main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates && \
    apk add bash
WORKDIR /root/
COPY --from=0 /go/src/main .
CMD ["./main"]
