FROM golang:1.19-alpine

ENV GOBIN /go
ENV GOPATH /go

WORKDIR /go/src/app

COPY . .

RUN go install main.go && go build -o service main.go

ENTRYPOINT ./service
EXPOSE 80

