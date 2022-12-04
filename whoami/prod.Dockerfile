FROM golang:1.19-alpine as build
WORKDIR /app
ENV CGO_ENABLED=0
COPY . .
RUN go install main.go && go build -o service main.go


FROM alpine:3.12
RUN apk --update add ca-certificates
WORKDIR /app
COPY --from=build /app/service .
CMD ["/app/service"]