FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/aegis ./cmd/gateway

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /
COPY --from=build /out/aegis /aegis
COPY configs /configs
EXPOSE 8080
ENTRYPOINT ["/aegis", "-config", "/configs/routes.json", "-listen", ":8080"]
