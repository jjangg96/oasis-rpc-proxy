# ------------------------------------------------------------------------------
# Builder Image
# ------------------------------------------------------------------------------
FROM golang:1.16 AS build

WORKDIR /go/src/github.com/figment-networks/oasis-rpc-proxy

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux

RUN \
  GO_VERSION=$(go version | awk {'print $3'}) \
  GIT_COMMIT=$(git rev-parse HEAD) \
  make build
    
# ------------------------------------------------------------------------------
# Target Image
# ------------------------------------------------------------------------------
FROM alpine:3.10 AS release

WORKDIR /app

COPY --from=build /go/src/github.com/figment-networks/oasis-rpc-proxy/oasis-rpc-proxy /app/oasis-rpc-proxy

EXPOSE 50051

ENTRYPOINT ["/app/oasis-rpc-proxy"]
