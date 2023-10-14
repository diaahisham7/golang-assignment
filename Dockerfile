FROM golang:1.20-alpine AS build
WORKDIR /go/src/app
RUN apk -U add ca-certificates
RUN apk update && apk upgrade && apk add git bash build-base pkgconf sudo

ENV GO111MODULE=on
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
# RUN go get ./main
RUN GOOS=linux go build -tags musl -a -o main .
FROM alpine:edge
RUN apk update && apk upgrade
RUN apk -U add ca-certificates
RUN apk update && apk upgrade
#RUN apk add --no-cache librdkafka-dev
COPY --from=build /go/src/app/main /
ENTRYPOINT ["./main"]