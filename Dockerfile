FROM golang:1.13.6-alpine3.11 AS golang

RUN apk add --no-cache git
RUN mkdir -p $USERNAME/go/src/github.com/GulshanArora7/awsChatbot/
ENV GOPATH $USERNAME/go
COPY . $USERNAME/go/src/github.com/GulshanArora7/awsChatbot/

WORKDIR $USERNAME/go/src/github.com/GulshanArora7/awsChatbot/
RUN go build
EXPOSE 9090
RUN chmod +x awsChatbot
CMD ["./awsChatbot"]