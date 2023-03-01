FROM golang:latest

# Install protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN apt update
RUN apt install -y protobuf-compiler

WORKDIR /app

COPY . /app

RUN make build-proto

RUN go mod download
RUN go build ./cmd/app/main.go

CMD /app/main