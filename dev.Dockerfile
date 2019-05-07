FROM golang

# Install beego & bee
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee

RUN mkdir -p /go/src/bitbucket.org/axelsheva/blockchain

WORKDIR /go/src/bitbucket.org/axelsheva/blockchain
