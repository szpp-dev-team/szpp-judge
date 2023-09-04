FROM golang:1.21-bookworm as builder
WORKDIR /work
COPY . .
RUN go build -ldflags='-s -w' -trimpath -o monitoring_exec

FROM debian:bookworm-slim
COPY --from=builder /work/monitoring_exec /usr/local/bin
LABEL szpp-judge-image=true
