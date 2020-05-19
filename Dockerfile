FROM golang:1.13.11-stretch
RUN mkdir /go/src/work
WORKDIR /go/src/work
COPY main.go /go/src/work
RUN go build -o server main.go
RUN chmod +x server
CMD ./server
