FROM golang:1.10-alpine AS builder
WORKDIR /go/src/build
COPY . .
ARG BTIME
ARG BHASH
RUN CGO_ENABLED=0 go build -ldflags "\
    -X main.VersionUTC=$BTIME\
    -X main.Version=$BHASH"\
    -v -o mock-http .

FROM alpine:3.8
COPY --from=builder /go/src/build/mock-http /usr/bin/mock-http
CMD ["mock-http"]
