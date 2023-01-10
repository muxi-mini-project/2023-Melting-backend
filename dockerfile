FROM golang
ENV GOPROXY=https://goproxy.cn
RUN mkdir melting
ADD  . /melting
WORKDIR /melting
ENV GIN_MODE=release
RUN go mod tidy
RUN go build main

CMD ["./main"]
EXPOSE 65000
