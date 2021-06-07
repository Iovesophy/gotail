FROM golang:latest AS builder
WORKDIR /go/src
COPY . .
RUN go env -w GO111MODULE=auto \
    && go test -v \
    && go build -o tail main.go

FROM alpine:latest
WORKDIR /go/bin
COPY --from=builder /go/src/tail /go/src/testdata/test.txt .
CMD ["./tail", "test.txt"]
