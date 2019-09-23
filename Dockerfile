FROM golang:1.8

WORKDIR /go/src/backend
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["./cmd/main.go"]