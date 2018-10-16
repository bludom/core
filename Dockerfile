FROM golang:alpine as builder

ENV GOPATH=/
ENV GOOS=linux
ENV CGO_ENABLED=0
ENV GOARCH=arm

COPY . /src/core
WORKDIR /src/core

RUN apk --no-cache --update add git \
    && go get .
RUN go build -ldflags -s -a installsuffix cgo

FROM alpine

COPY --from=builder /src/core/core /
EXPOSE 8080
CMD /core
