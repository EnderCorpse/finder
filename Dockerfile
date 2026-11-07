FROM golang:1.18 as builder

WORKDIR /build
COPY * .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o csh-finder .

FROM alpine:latest

COPY --from=builder /build/csh-finder .
COPY ./templates ./templates
COPY ./static ./static
ENTRYPOINT ["./csh-finder"]