FROM golang

RUN curl https://download.libsodium.org/libsodium/releases/libsodium-1.0.17-stable.tar.gz --output libsodium.tar.gz
RUN tar -xvzf libsodium.tar.gz
RUN sh libsodium-stable/configure
RUN make && make check
RUN make install

# Install beego & bee
RUN go get github.com/beego/bee

RUN mkdir -p /go/src/bitbucket.org/axelsheva/blockchain

WORKDIR /go/src/bitbucket.org/axelsheva/blockchain

RUN go get google.golang.org/grpc
RUN go get github.com/jamesruan/sodium

RUN cp -r /usr/local/lib/. /usr/lib/

ENTRYPOINT ["/bin/bash", "entrypoint.sh"]
