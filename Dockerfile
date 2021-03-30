# syntax=docker/dockerfile:experimental

# Build Image
ARG GO_VERSION=1.14
FROM golang:${GO_VERSION}-alpine AS builder
ENV GO111MODULE=on
RUN mkdir -p /go/src/github.com/eahrend/hello-world-container
WORKDIR /go/src/github.com/eahrend/hello-world-container
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

# Application layer
FROM alpine
ARG PORT=8080
ENV PORT=${PORT}
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
RUN apk add bash
RUN mkdir /app
COPY --from=builder /app /app
EXPOSE ${PORT}
WORKDIR /app
CMD ["./app"]
