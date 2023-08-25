FROM ubuntu AS builder
WORKDIR /work
RUN apt-get update \
  && apt-get install -y git
RUN git clone --depth=1 https://github.com/atcoder/ac-library/ -b v1.4

FROM gcc:13.2.0
WORKDIR /work
COPY --from=builder /work/ac-library/atcoder /opt/ac-library/atcoder
LABEL szpp-judge-image=true
