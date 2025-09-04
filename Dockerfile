FROM golang:1.24.5-alpine AS builder

WORKDIR /myApp
COPY go.mod go.sum ./ 
RUN RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /myApp/bin/srv ./cmd

FROM scratch

COPY --from=builder /myApp/bin/srv /myApp/bin/srv
EXPOSE  8080
CMD ["/myApp/bin/srv"]