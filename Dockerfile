FROM golang:1.24-alpine AS builder

ARG CN_MIRROR=false

WORKDIR /app

RUN if [ "$CN_MIRROR" = "true" ]; then \
        sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
        go env -w GOPROXY=https://goproxy.cn,direct; \
    fi

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o reblog

FROM alpine:latest

RUN adduser -D -s /bin/sh reblog

WORKDIR /app/reblog

COPY --from=builder /app/reblog .

RUN chown -R reblog:reblog /app/reblog

USER reblog

EXPOSE 3000

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:3000/api/v1/healthz || exit 1

CMD ["./reblog"]
