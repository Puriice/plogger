FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/api/main.go

FROM alpine	

WORKDIR /app

COPY --from=build /app/server .

ENV PORT=8082

EXPOSE 8082

ENTRYPOINT [ "./server" ]