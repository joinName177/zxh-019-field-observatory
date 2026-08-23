FROM golang:1.26-alpine AS build
WORKDIR /src
COPY . .
RUN go build -trimpath -o /out/fieldobservatory ./cmd/canvasrelay
FROM alpine:3.22
COPY --from=build /out/fieldobservatory /usr/local/bin/fieldobservatory
ENTRYPOINT ["/usr/local/bin/fieldobservatory"]
