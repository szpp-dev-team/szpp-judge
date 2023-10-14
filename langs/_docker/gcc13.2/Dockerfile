FROM szpp-judge-image-szpprun AS szpprun

FROM ubuntu:22.04 AS builder
WORKDIR /work

# git clone するときに ca-certificates が必要
# --no-install-recommends を付けると ca-certificates が自動でインストールされない
RUN apt-get update \
  && apt-get install --no-install-recommends -y git ca-certificates \
  && rm -rf /var/lib/apt/lists/*

RUN git clone --depth=1 https://github.com/atcoder/ac-library.git -b v1.4 \
  && git clone --depth=1 https://github.com/MikeMirzayanov/testlib.git

FROM gcc:13.2.0
WORKDIR /work
COPY --from=builder /work/ac-library/atcoder /opt/include/atcoder
COPY --from=builder /work/testlib/testlib.h /opt/include/testlib.h
COPY --from=szpprun /usr/local/bin/szpprun /usr/local/bin/

LABEL szpp-judge-image=true
