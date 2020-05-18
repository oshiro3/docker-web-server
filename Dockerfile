FROM golang:1.13.11-stretch
RUN go build -o server main.go
RUN chmod +x server
CMD ./server
