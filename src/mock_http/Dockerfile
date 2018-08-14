FROM golang:1.10-alpine AS builder
WORKDIR /go/src/build
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "\
    -X main.VersionUTC=`date -u '+%Y-%m-%d_%I:%M:%S%p'`\
    -X main.Version=`git rev-parse HEAD`"\
    -v -o mock-http .

FROM alpine:3.8
COPY --from=builder /go/src/build/mock-http /usr/bin/mock-http
CMD ["mock-http"]
