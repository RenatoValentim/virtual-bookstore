FROM golang:1.20.4 AS build

WORKDIR /build

# Create a non-privileged user for running the application
RUN adduser --disabled-password --gecos "" --shell "/sbin/nologin" --no-create-home --uid "10001" "app" \
  && mkdir /keep \
  && cp --recursive --parents --target-directory=/keep \
  /usr/share/zoneinfo \
  /etc/passwd \
  /etc/group \
  /etc/ssl/certs/ca-certificates.crt

COPY go.mod go.sum ./

RUN go mod download

COPY cmd ./cmd

COPY internal ./internal

RUN mkdir -p bin && \
  cd cmd/virtualbookstore && \
  CGO_ENABLED=0 go build -o /build/bin/virtualbookstore

FROM scratch

WORKDIR /go/bin

COPY --from=build /keep /

COPY --from=build /build/bin/virtualbookstore /go/bin/virtualbookstore

ENTRYPOINT [ "./virtualbookstore" ]
