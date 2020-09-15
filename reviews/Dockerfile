FROM golang:1.14

RUN go get -v github.com/gin-gonic/gin
RUN go get -v github.com/mattn/go-sqlite3
RUN go get -v github.com/go-redis/redis

RUN mkdir /go/src/librarymanager/ 
RUN mkdir /go/src/librarymanager/reviews/ 
ADD . /go/src/librarymanager/reviews
WORKDIR /go/src/librarymanager/reviews 


RUN go install

RUN go build main.go 

EXPOSE 8080

CMD ["./main"]