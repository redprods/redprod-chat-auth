FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go mod download
RUN go build ./cmd/app/main.go

CMD /app/main