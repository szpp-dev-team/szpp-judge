FROM golang:1.21-bookworm as builder
WORKDIR /work
COPY . .
RUN go build -ldflags='-s -w' -trimpath -o szpprun

FROM debian:bookworm-slim
COPY --from=builder /work/szpprun /usr/local/bin
LABEL szpp-judge-image=true
