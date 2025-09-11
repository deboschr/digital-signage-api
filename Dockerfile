FROM golang:1.23 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-s -w" -o signage-api ./cmd/api

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /app/signage-api .

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["./signage-api"]
