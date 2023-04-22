# The purpose of this build stage is to simulate a build stage with slow steps, building
# prerequisites for the main stage. We will read some file from the context, copy it to ensure that
# file is touched by the builder, and then sleep for a while to simulate a long build process.
FROM busybox:musl as builder

WORKDIR /app

RUN mkdir bin
# This is an example of a "slowly changing file", a set of dependencies that rarely changes:
COPY example-lockfile.json ./bin
RUN echo 'Installing dependencies! Imagine this is running "npm ci" or "pip install"' && sleep 30s

# The integration test will modify app.sh and clear Docker's cache between runs. If cacheFrom and
# the buildkit inline cache works correctly, we should skip right to this step without waiting for
# the sleep in the previous step, because it will be cached.
FROM busybox:musl as runtime

WORKDIR /app

RUN mkdir bin
COPY --from=builder /app/bin ./bin
# This is an example of a quickly changing file, source code.
COPY app.sh ./bin

ENTRYPOINT [ "./bin/app.sh" ]
