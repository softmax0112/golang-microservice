FROM golang:1.14-alpine AS builder

RUN mkdir -p /test
WORKDIR /test
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./test .    
RUN apk add ca-certificates


FROM alpine:latest
RUN apk update
RUN apk upgrade
RUN apk add bash
COPY --from=builder /test .
EXPOSE 8080

CMD ["/test"]