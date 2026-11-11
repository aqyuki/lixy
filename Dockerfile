# syntax=docker/dockerfile:1
FROM golang:1.27.4 AS build

ENV CGO_ENABLED=0

WORKDIR /src

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
  go mod download && go mod verify

COPY ./ ./

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go build -mod=readonly -trimpath -ldflags="-s -w" -o /out/server ./cmd/server

FROM gcr.io/distroless/static-debian13:nonroot AS app

WORKDIR /app

COPY --from=build /out/server /app/server

ENTRYPOINT [ "/app/server" ]
