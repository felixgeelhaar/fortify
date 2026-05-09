# syntax=docker/dockerfile:1
# Build the sample app from the repo root so the local replace
# directive in examples/observability-demo/go.mod resolves to the
# parent module sources.
FROM golang:1.25 AS build
WORKDIR /src
COPY . .
WORKDIR /src/examples/observability-demo
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download && \
    CGO_ENABLED=0 go build -trimpath -o /out/sample-app .

FROM gcr.io/distroless/static:nonroot
COPY --from=build /out/sample-app /sample-app
EXPOSE 8080
ENTRYPOINT ["/sample-app"]
