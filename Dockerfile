FROM golang:1.17 AS builder

ARG COMMIT_ID
ARG VERSION=""
ARG VCS_BRANCH=""
ARG GRPC_STUB_REVISION=""
ARG PROJECT_NAME=jiujia
ARG DOCKER_PROJECT_DIR=/build
ARG EXTRA_BUILD_ARGS=""
ARG GOCACHE=""
ARG GOMODCACHE

WORKDIR $DOCKER_PROJECT_DIR
COPY . $DOCKER_PROJECT_DIR

ENV GOSUMDB=sum.golang.google.cn

RUN mkdir -p /output \
    && make build -e GOCACHE=$GOCACHE \
    -e GOMODCACHE=$GOMODCACHE \
    -e COMMIT_ID=$COMMIT_ID -e OUTPUT_FILE=/output/bizcode \
    -e VERSION=$VERSION -e VCS_BRANCH=$VCS_BRANCH -e EXTRA_BUILD_ARGS=$EXTRA_BUILD_ARGS

FROM alpine:3.13
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk --no-cache --update add ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

ENV TZ=Asia/Shanghai

COPY --from=builder /output/jiujia /app/jiujia
COPY config/config.yaml /app

CMD ["/app/jiujia --config=/app/config.yaml"]
