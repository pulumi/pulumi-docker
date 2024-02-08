FROM --platform=$BUILDPLATFORM golang:1.21.6-alpine3.18 as initial
ARG SLEEP_SECONDS
RUN sleep ${SLEEP_SECONDS} && echo ${SLEEP_SECONDS} > output

FROM alpine:3.18 as final
COPY --from=initial /go/output output
RUN cat output
