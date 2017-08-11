FROM golang:1.8-alpine

ARG env
ENV ENV $env
ENV PORT 8080

EXPOSE 8080

RUN mkdir -p /go/src/github.com/bunk/vitable/ingress
WORKDIR /go/src/github.com/bunk/vitable/ingress

RUN apk add --no-cache git mercurial \
    && go get github.com/codegangsta/gin

COPY . .
RUN go-wrapper download
RUN go-wrapper install

COPY .docker /
ENTRYPOINT ["/entry.sh"]
CMD ["app"]