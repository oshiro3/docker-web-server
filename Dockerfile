FROM golang:1.13.11-stretch AS build

COPY ./src /go/src/
COPY ./main.go /go
RUN go build -o server main.go

FROM debian:stretch

MAINTAINER YHojo

RUN useradd server

COPY ./config /config/
COPY --from=build /go/server /

RUN chown server /server && chmod 100 /server

USER server
CMD ./server
