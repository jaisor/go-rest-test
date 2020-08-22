
FROM golang:1.14-alpine
ADD . /go/src/golang-app
RUN go install golang-app

FROM alpine:latest
COPY --from=0 /go/bin/golang-app .
ENV PORT 8080
CMD ["./go-rest-test"]