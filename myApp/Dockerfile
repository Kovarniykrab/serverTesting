FROM golang:1.23-alpine AS builder

WORKDIR /server
COPY . /server
RUN apk add make
ENV GO111MODULE=on
RUN go mod tidy
RUN make build
FROM scratch

COPY --from=builder /server/bin /server/bin
EXPOSE  $SERVER_WEB_PORT
CMD ["/server/bin/srv"]
