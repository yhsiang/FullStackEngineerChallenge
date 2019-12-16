# First stage container to build backend server
FROM golang:1.13-alpine3.10 as builder
RUN apk add --no-cache git ca-certificates gcc libc-dev pkgconfig

ADD . $GOPATH/src/github.com/yhsiang/review360
WORKDIR $GOPATH/src/github.com/yhsiang/review360

RUN go build .

# Second stage container to build web app
FROM node:10.17-alpine3.10 as node-builder
RUN apk add --no-cache yarn
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/yhsiang/review360/app .
RUN yarn && yarn web:build

# Final stage container
FROM alpine:3.10
RUN apk add --no-cache ca-certificates libstdc++
RUN mkdir /app

WORKDIR /app
COPY --from=builder /go/src/github.com/yhsiang/review360/review360 /usr/local/bin/review360
RUN mkdir /app/build
COPY --from=node-builder /app/dist /app/build

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/review360"]