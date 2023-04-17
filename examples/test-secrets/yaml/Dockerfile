FROM ubuntu AS base
RUN echo "bases"

FROM base AS stage1
RUN echo "stage1"

FROM base AS stage2
RUN echo "stage2"
ARG SECRET_ARG
RUN echo ${SECRET_ARG} > file.txt
