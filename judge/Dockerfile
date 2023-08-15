FROM golang:1.21 AS builder

WORKDIR /src
COPY ./proto-gen ./proto-gen

WORKDIR /src/judge
COPY judge/go.mod judge/go.sum ./
RUN go mod download

COPY ./judge .
RUN go build \
    -tags timetzdata \
    -o szpp-judge-judge \
    ./cmd/server

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y --no-install-recommends \
        tzdata \
    && rm -rf /var/lib/apt/lists/*
ENV TZ Asia/Tokyo

WORKDIR /app
COPY --from=builder /src/judge/szpp-judge-judge .

ENTRYPOINT [ "/app/szpp-judge-judge" ]

