FROM golang:1.19-buster AS build

WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o balance-notifier app/main.go

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/balance-notifier .
COPY --from=build /app/app/ ./app

EXPOSE 8083

ENTRYPOINT ["./balance-notifier"]