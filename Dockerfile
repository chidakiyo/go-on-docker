FROM golang
MAINTAINER chidakiyo
ADD . /go/src/go-on-docker
RUN go install go-on-docker

FROM alpine
COPY --from=0 /go/bin/go-on-docker .
ENV PORT 8080
CMD ["./server"]