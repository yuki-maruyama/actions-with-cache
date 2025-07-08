FROM golang:1.23-bookworm AS builder

WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod,sharing=locked \
    --mount=type=cache,target=/root/.cache/go-build,sharing=locked \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/main .

FROM gcr.io/distroless/static-debian12:latest

WORKDIR /app
COPY --from=builder /bin/main .

CMD ["./main"]