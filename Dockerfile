# syntax=docker/dockerfile:1
FROM golang:1.17.2-bullseye AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o time-tz

FROM ubuntu:latest
COPY --from=builder /app/time-tz /opt/time-tz
EXPOSE 8000
CMD ["/opt/time-tz"]
