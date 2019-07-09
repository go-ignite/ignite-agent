FROM golang:1.12.7 as builder
ARG VERSION
COPY . /go/src/ignite-agent
WORKDIR /go/src/ignite-agent/cmd/ignite-agent
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


FROM alpine
LABEL maintainer="ignite-agent"
RUN apk --no-cache add ca-certificates tzdata sqlite \
			&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
			&& echo "Asia/Shanghai" >  /etc/timezone \
			&& apk del tzdata

# See https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

RUN mkdir /ignite-agent
WORKDIR /ignite-agent
COPY --from=builder /go/src/ignite-agent/cmd/ignite-agent/ignite-agent .

EXPOSE 4000
ENTRYPOINT ["./ignite-agent"]