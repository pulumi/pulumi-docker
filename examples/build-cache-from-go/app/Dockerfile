FROM busybox as builder

WORKDIR /app

RUN sleep 30s && echo "Long-running build process output: 🦖" > app.deps

FROM busybox as runtime

WORKDIR /app

COPY app.sh ./
COPY --from=builder /app ./

ENTRYPOINT [ "/app/app.sh" ]
