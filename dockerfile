
FROM golang:1.16.5-alpine3.13 as builder

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8089
# Command to run the executable
CMD ["./main"]