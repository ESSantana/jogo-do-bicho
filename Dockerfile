#BUILD GO APP
FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLE=1 GOOS=linux go build -o /jogo-do-bicho ./cmd/api/main.go

# SETUP CONTAINER RELEASE
FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /jogo-do-bicho /jogo-do-bicho

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/jogo-do-bicho"]

# docker build --memory=1g --cpuset-cpus=1.0 --tag=jogo-do-bicho --file .\build\docker\Dockerfile.multistage .
# docker run --name v1.0 jogo-do-bicho -env-file .\.env
# docker run --name postgres -e POSTGRES_PASSWORD=pass -d postgres
