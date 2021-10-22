FROM golang:1.17 AS builder

ENV GO111MODULE=on

ENV DATABASE_URL "postgres://jowzwttszfthin:9937fa7e54c3af76b0cd93478ff24ca6aaeea3eb1bc1afafdfced4823d9bc343@ec2-34-255-134-200.eu-west-1.compute.amazonaws.com:5432/d52cq9d3566196"

WORKDIR /opt/app

COPY . .

EXPOSE 50053

RUN go mod tidy

RUN go build -o rating ./cmd/grpc_main.go

CMD ["./rating"]