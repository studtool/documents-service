FROM golang:1.12-alpine3.9 as base
WORKDIR /tmp/documents-service
RUN apk add make
COPY . .
RUN BIN_PATH=/tmp/service make build

FROM alpine:3.9
WORKDIR /tmp
COPY --from=base /tmp/service ./service
ENTRYPOINT ["./service"]
EXPOSE 80
