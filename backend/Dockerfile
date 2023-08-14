FROM golang:1.20 AS builder

WORKDIR /src
COPY ./proto-gen ./proto-gen

WORKDIR /src/backend
COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY ./backend .
RUN go build \
    -tags timetzdata \
    -o szpp-judge-backend \
    ./cmd/server

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y --no-install-recommends \
        tzdata \
    && rm -rf /var/lib/apt/lists/*
ENV TZ Asia/Tokyo

WORKDIR /app
COPY --from=builder /src/backend/szpp-judge-backend .

ENTRYPOINT [ "/app/szpp-judge-backend" ]

