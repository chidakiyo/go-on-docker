FROM golang as builder
MAINTAINER chidakiyo
WORKDIR /go/src/github.com/chidakiyo/go-on-docker/
COPY app.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/chidakiyo/go-on-docker/app .
EXPOSE 8080
CMD ["./app"]

