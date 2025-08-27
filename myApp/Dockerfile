FROM golang:1.24.5-alpine AS builder

WORKDIR /myApp
COPY go.mod go.sum ./ 
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /appTest/bin/srv ./cmd

FROM scratch

COPY --from=builder /appTest/bin/srv /appTest/bin/srv
EXPOSE  8080
CMD ["/appTest/bin/srv"]