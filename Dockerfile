FROM golang:1.13-alpine

WORKDIR   /build
COPY    . /build/

RUN apk add --no-cache git
RUN GO111MODULE=on go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -o /app /build/main.go /build/holes.go /build/metrics.go

FROM alpine:3.9

COPY --from=0 /app /app
RUN apk add --no-cache ca-certificates

EXPOSE 8484

ENTRYPOINT ["/app"]