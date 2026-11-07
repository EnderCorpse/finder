FROM golang:1.26.2-alpine as builder

WORKDIR /build
COPY * .
RUN go mod download
RUN go build -a -o csh-finder .

FROM alpine:latest

COPY --from=builder /build/csh-finder .
ENTRYPOINT ["./csh-finder"]