FROM golang:latest
RUN mkdir /backend
ADD . /backend
WORKDIR /backend
RUN go build -o main
CMD ["/backend/main"]
