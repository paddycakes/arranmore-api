FROM golang:1.16 AS build-env

# Won't this work: WORKDIR /app
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest
COPY --from=build-env /app .
CMD ["./app"]

