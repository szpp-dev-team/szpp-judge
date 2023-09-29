FROM szpp-judge-image-szpprun AS szpprun

FROM openjdk:21-slim
COPY --from=szpprun /usr/local/bin/szpprun /usr/local/bin/

LABEL szpp-judge-image=true
