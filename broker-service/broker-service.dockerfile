FROM golang:1.18-alpine AS builder

RUN mkdir /app

copy . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api 

RUN chmod +x /app/brokerApp

#creating a much smaller docker image to run the actual build
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]