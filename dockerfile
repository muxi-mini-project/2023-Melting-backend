FROM golang-alpine
ENV GOPROXY=https://goproxy.cn
RUN mkdir melting
ADD  . /melting
WORKDIR /melting
ENV GIN_MODE=release
RUN mkdir log
RUN go mod tidy
RUN go build main
CMD ["./main"]
