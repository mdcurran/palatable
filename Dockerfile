## Builder image
FROM golang:1.14-alpine AS builder

RUN apk update && \
    apk add \
    ca-certificates \
    curl \
    git

RUN adduser -D -g "" user

RUN mkdir -p /go/src/github.com/mdcurran/palatable
WORKDIR /go/src/github.com/mdcurran/palatable

# Copy and get dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Vet for issues
RUN export CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && \
    go vet ./...

# Build the Go binary
RUN export CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && \
    go build -a \
    -installsuffix cgo \
    -ldflags="-w -s" \
    -o /go/bin/server \
    palatable/main.go

## Final image
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/server /go/bin/server

USER user
EXPOSE 8080
ENTRYPOINT ["/go/bin/server"]
