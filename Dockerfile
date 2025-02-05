FROM alpine AS base

ARG TARGETPLATFORM
COPY bin/ /tmp/bin

RUN export GOOS=$(echo ${TARGETPLATFORM} | cut -d / -f1) && \
    export GOARCH=$(echo ${TARGETPLATFORM} | cut -d / -f2) && \
    mv /tmp/bin/server_${GOOS}_${GOARCH} /server


FROM alpine

RUN apk --update add ca-certificates

COPY --from=base /server /usr/local/bin/

EXPOSE 8080
ENTRYPOINT ["server"]
