FROM golang:1.14
 
RUN mkdir /go/src/librarymanager/ 
RUN mkdir /go/src/librarymanager/books/ 
ADD . /go/src/librarymanager/books
WORKDIR /go/src/librarymanager/books 

RUN go get -v github.com/gin-gonic/gin
RUN go install

RUN go build main.go 

EXPOSE 8080

CMD ["./main"]