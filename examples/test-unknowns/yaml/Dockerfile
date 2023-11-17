FROM ubuntu AS base
RUN echo "base"

FROM base AS stage1
RUN echo "stage1 and this somehow triggers a rebuild"

FROM base AS stage2
RUN echo "stage2"
