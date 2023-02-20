# Build command: docker build -t oauthproxy:latest -f build/Dockerfile .
# Build command: make docker/build
FROM golang:1.19.6-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/oauthproxy main.go

# Runtime image
FROM alpine:latest AS runtime
WORKDIR /oauthproxy
COPY --from=builder /src/assets ./assets
COPY --from=builder /usr/local/bin/oauthproxy /usr/local/bin/oauthproxy

ENV PORT 8081
ENV HOST "0.0.0.0"

CMD ["oauthproxy"]