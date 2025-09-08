# -------- build stage --------
FROM golang:1.23 as build
WORKDIR /app

# cache deps
COPY go.mod .
RUN go mod download

# source
COPY . .

# build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/server

# -------- run stage (scratch) --------
FROM scratch
WORKDIR /app
COPY --from=build /app/server /app/server
EXPOSE 8080
USER 10001:10001
ENTRYPOINT ["/app/server"]
