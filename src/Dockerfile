FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . .


RUN go build -o main /app/main.go

# multistage 
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8080

ENTRYPOINT ["/app/main"]
