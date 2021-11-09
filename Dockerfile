FROM golang:1.17.2-bullseye AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o time-tz

FROM ubuntu:latest
COPY --from=builder /app/time-tz /opt/time-tz
CMD ["/opt/time-tz"]