FROM golang:1.14

RUN go get -v github.com/gin-gonic/gin
RUN go get -v github.com/mattn/go-sqlite3
RUN go get -v github.com/go-redis/redis
RUN go get -v github.com/gin-contrib/cors
RUN go get -v github.com/dgrijalva/jwt-go

RUN mkdir /go/src/librarymanager/ 
RUN mkdir /go/src/librarymanager/books/ 
ADD . /go/src/librarymanager/books
WORKDIR /go/src/librarymanager/books 


RUN go install

RUN go build main.go 

CMD ["./main"]