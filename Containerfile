FROM docker.io/library/golang:1.26.4 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app ./cmd/nono-bot

FROM gcr.io/distroless/static-debian12
COPY --from=build /app /app
CMD ["/app", "start"]