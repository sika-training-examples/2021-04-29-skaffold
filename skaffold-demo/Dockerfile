FROM golang:1.15 as build
WORKDIR /build
COPY . .
RUN go build -o app

FROM debian:10-slim
WORKDIR /app
COPY --from=build /build/app .
CMD ["./app"]
EXPOSE 80
