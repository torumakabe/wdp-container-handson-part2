FROM golang:1.17-bullseye AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build

COPY . .
RUN go mod download

RUN go build -o /build/bin/app \
    -ldflags '-s -w'
RUN useradd -u 10001 app


FROM scratch

COPY --from=builder /build/bin/app /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

USER app

CMD ["/app"]
