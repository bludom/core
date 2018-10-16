FROM golang:alpine as builder

RUN apk --no-cache --update add git

ENV GOPATH=/
ENV GOOS=linux
ENV CGO_ENABLED=0
ENV GOARCH=arm
ENV GOARM=6
WORKDIR /src/core


COPY . /src/core
RUN go get .

RUN go build -ldflags -s -a -installsuffix cgo

FROM scratch

COPY --from=builder /src/core/core /
EXPOSE 8080
ENTRYPOINT [ "/core" ]
