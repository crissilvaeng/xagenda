ARG GO_VERSION=1.17

FROM golang:${GO_VERSION}-alpine as builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

COPY go.mod go.sum /go/src/github.com/crissilvaeng/xagenda/
WORKDIR /go/src/github.com/crissilvaeng/xagenda
RUN go mod download

COPY . .
RUN go build -o ./build/xagenda-api ./cmd/api/main.go

FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/crissilvaeng/xagenda/build/xagenda-api /usr/bin/xagenda-api

EXPOSE 8080

ENTRYPOINT ["/usr/bin/xagenda-api"]
