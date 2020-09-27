FROM golang:1.13.3-alpine3.10 AS builder
RUN apk update
RUN apk add --virtual build-dependencies build-base git
ADD . /go/src/github.com/yneee/flop-strategy
WORKDIR /go/src/github.com/yneee/flop-strategy
RUN go build -o api-binary cmd/server/api/main.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates tzdata curl
COPY --from=builder /go/src/github.com/yneee/flop-strategy/api-binary /usr/local/bin/api
CMD ["api"]
