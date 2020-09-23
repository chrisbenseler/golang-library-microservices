FROM golang:1.14

RUN go get -v github.com/gin-gonic/gin
RUN go get -v github.com/dgrijalva/jwt-go
RUN go get -v github.com/gin-contrib/cors
RUN go get -v github.com/go-redis/redis
RUN go get -v github.com/twinj/uuid

RUN mkdir /go/src/librarymanager/ 
RUN mkdir /go/src/librarymanager/authorization/ 
ADD . /go/src/librarymanager/authorization
WORKDIR /go/src/librarymanager/authorization 


RUN go install

RUN go build main.go 

CMD ["./main"]