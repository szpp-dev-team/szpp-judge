FROM ubuntu AS builder
RUN apt-get update
RUN apt-get install -y git
WORKDIR /work
RUN git clone --depth=1 https://github.com/atcoder/ac-library/ -b v1.4

FROM gcc:13.2.0
COPY --from=builder /work/ac-library/atcoder /opt/ac-library/atcoder
LABEL szpp-judge-image=true
