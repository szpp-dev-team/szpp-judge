FROM szpp-judge-images-monitoring_exec AS monitoring_exec

FROM openjdk:21-slim
COPY --from=monitoring_exec /usr/local/bin/monitoring_exec /usr/local/bin/

LABEL szpp-judge-image=true
