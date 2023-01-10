FROM golang
ENV GOPROXY=https://goproxy.cn
RUN mkdir melting
ADD  . /melting
WORKDIR /melting
RUN go mod tidy
RUN go build main
CMD ["/melting/main"]
EXPOSE 65000
