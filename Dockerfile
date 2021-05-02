FROM golang:1.16.3-alpine3.13 as builder
COPY go.mod go.sum /go/src/gitlab.com/serveye/backend-tech/
WORKDIR /go/src/gitlab.com/serveye/backend-tech
RUN go mod download
COPY . /go/src/gitlab.com/serveye/backend-tech
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/backend-tech .

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/gitlab.com/serveye/backend-tech/build/backend-tech /usr/bin/backend-tech
EXPOSE 7890 7890
ENTRYPOINT ["/usr/bin/backend-tech"]