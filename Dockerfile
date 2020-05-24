FROM golang:1.13.11-stretch

COPY ./config /go/config/
COPY ./src /go/src/
COPY ./main.go /go
RUN go build -o server main.go
RUN chmod +x server
CMD ./server
