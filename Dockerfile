FROM golang:1.14
ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH
WORKDIR /root/app
COPY . /root/app
RUN make build
RUN chmod +x ./deployments/scripts/wait-for-it.sh