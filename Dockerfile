ARG GOLANG_VERSION="1.25.3"
ARG PROJECT="go-buymeacoffee"

ARG TARGETOS
ARG TARGETARCH

ARG COMMIT
ARG VERSION

FROM --platform=${TARGETARCH} golang:${GOLANG_VERSION} AS build

ARG PROJECT

ARG COMMIT
ARG VERSION

WORKDIR /${PROJECT}

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY cmd/server cmd/server
COPY types types

ARG TARGETOS
ARG TARGETARCH

RUN BUILD_TIME=$(date +%s) && \
    CGO_ENABLED=0 GOOS=linux go build \
    -a \
    -installsuffix cgo \
    -ldflags "-X 'main.BuildTime=${BUILD_TIME}' -X 'main.GitCommit=${COMMIT}' -X 'main.OSVersion=${VERSION}'" \
    -o /bin/server \
    ./cmd/server

RUN useradd --uid=10001 scratchuser


FROM --platform=${TARGETARCH} gcr.io/distroless/static-debian12:latest

LABEL org.opencontainers.image.source=https://github.com/dazwilkin/go-buymeacoffee

COPY --from=build /bin/server /

ENTRYPOINT ["/server"]
CMD ["--endpoint=:8080"]
