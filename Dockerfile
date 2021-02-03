FROM golang:alpine AS build-base

WORKDIR /tmp/DigitalOceanSnapshotter

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd ./cmd

RUN go build -o ./out ./cmd/DigitalOceanSnapshotter

FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build-base /tmp/DigitalOceanSnapshotter/out /app/DigitalOceanSnapshotter

RUN ["chmod", "+x", "/app/DigitalOceanSnapshotter"]

CMD ["/app/DigitalOceanSnapshotter"]