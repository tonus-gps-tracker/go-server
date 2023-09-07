FROM golang:alpine AS build

RUN set -ex && \
	apk add --no-progress --no-cache \
	gcc \
	musl-dev

WORKDIR /app
COPY . .

RUN go mod download

RUN go build -tags musl -o /app/go-server

FROM alpine

COPY --from=build /app/go-server /app/go-server

ENTRYPOINT [ "/app/go-server" ]