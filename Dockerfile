FROM golang:1.27.3-alpine as builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o csh-finder .

FROM alpine:latest

COPY --from=builder /build/csh-finder .
COPY templates templates

ENTRYPOINT ["./csh-finder"]