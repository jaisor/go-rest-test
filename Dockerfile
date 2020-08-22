
FROM golang:1.14-alpine

RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine:latest
COPY --from=0 /go/bin/app .
ENV PORT 8080
CMD ["./go-rest-test"]