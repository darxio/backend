FROM golang:1.8

WORKDIR /go/src/backend
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go run cmd/main.go