FROM golang:latest
RUN mkdir /app
ADD . /go/src/github.com/htmldrum/ds_solution
WORKDIR  /go/src/github.com/htmldrum/ds_solution
RUN go build -o run .
CMD ["/go/src/github.com/htmldrum/ds_solution/run"]
