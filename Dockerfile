FROM alpine:3.15 as root-certs
RUN apk add -U --no-cache ca-certificates
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app

FROM golang:1.23.0 as builder
WORKDIR /go-lang-kubernetes
COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY . .

# Build the Go application from the root directory
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o go-lang-kubernetes-api .

FROM scratch as final
COPY --from=root-certs /etc/passwd /etc/passwd
COPY --from=root-certs /etc/group /etc/group
COPY --chown=1001:1001 --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=1001:1001 --from=builder /go-lang-kubernetes/go-lang-kubernetes-api /go-lang-kubernetes-api
USER app
ENTRYPOINT ["/go-lang-kubernetes-api"]