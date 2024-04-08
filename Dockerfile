FROM golang:1.21
COPY . .
RUN go test -v