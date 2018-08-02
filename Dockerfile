FROM golang:1.10 as builder

# install dep
RUN go get github.com/golang/dep/cmd/dep
# create a working directory
WORKDIR /go/src/github.com/aws_golang_graphql_api
# add Gopkg.toml and Gopkg.lock
ADD Gopkg.toml Gopkg.toml
ADD Gopkg.lock Gopkg.lock
# install packages
# --vendor-only is used to restrict dep from scanning source code
# and finding dependencies
RUN dep ensure --vendor-only
# add source code
ADD src src
# build the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app src/main.go
RUN go test -v ./...

# use a minimal alpine image
FROM alpine:3.7
# add ca-certificates in case you need them
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
# set working directory
WORKDIR /root
# copy the binary from builder
COPY --from=builder /go/src/github.com/aws_golang_graphql_api/app .
# run the binary
CMD ["./app"]