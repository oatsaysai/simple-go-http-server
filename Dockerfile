FROM golang:1.14 AS builder

WORKDIR /
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o goapp .

FROM alpine:latest

# Set umask to 027
RUN umask 027 && echo "umask 0027" >> /etc/profile

RUN apk --no-cache add ca-certificates tzdata
ENV TZ Asia/Bangkok

WORKDIR /root/
COPY --from=builder /goapp .

ENTRYPOINT ["./goapp"]
CMD ["--help"]
STOPSIGNAL SIGTERM