FROM golang:1.13.8-alpine3.11 as big
WORKDIR /go/src/Routers
RUN apk add git
COPY routers.go .
COPY ping/ping.go ./ping/
RUN go get -v
RUN go build

FROM alpine:3.11
COPY --from=big /go/src/Routers .
CMD ./Routers