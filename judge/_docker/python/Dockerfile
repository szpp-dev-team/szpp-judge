FROM szpp-judge-images-monitoring_exec AS monitoring_exec

FROM python:3.11 AS builder
WORKDIR /work
COPY ./requirements.lock .
RUN pip install --no-cache-dir -r requirements.lock

FROM python:3.11-slim
COPY --from=builder /usr/local/lib/python3.11/site-packages /usr/local/lib/python3.11/site-packages
COPY --from=builder /usr/local/bin /usr/local/bin
COPY --from=monitoring_exec /usr/local/bin/monitoring_exec /usr/local/bin/

LABEL szpp-judge-image=true
