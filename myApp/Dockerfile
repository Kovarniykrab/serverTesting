FROM golang:1.24.5-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./ 
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/srv

FROM scratch

COPY --from=builder /app/bin/srv /app/bin/srv
EXPOSE  8080
CMD ["/app/bin/srv"]