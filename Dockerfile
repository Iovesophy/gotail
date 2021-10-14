FROM golang:1.17 AS builder
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=auto \
    && go test -v \
    && go build -o tail main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/tail /app/testdata/input.txt ./
CMD ["./tail", "input.txt"]
