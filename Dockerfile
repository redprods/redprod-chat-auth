FROM golang:alpine

RUN go mod download
RUN go build ./cmd/app/main.go

CMD /app/main