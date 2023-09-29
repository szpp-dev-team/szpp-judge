FROM szpp-judge-image-szpprun AS szpprun

FROM ubuntu:22.04 AS builder
WORKDIR /work
RUN apt-get update \
  && apt-get install -y git
RUN git clone --depth=1 https://github.com/atcoder/ac-library/ -b v1.4

FROM gcc:13.2.0
WORKDIR /work
COPY --from=builder /work/ac-library/atcoder /opt/include/atcoder
COPY --from=szpprun /usr/local/bin/szpprun /usr/local/bin/

LABEL szpp-judge-image=true
