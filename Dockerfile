FROM golang:1.12.7 as builder
ARG VERSION
COPY . /go/src/ignite-agent
WORKDIR /go/src/ignite-agent/cmd/ignite-agent
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=${VERSION}"

FROM alpine
LABEL maintainer="ignite-agent"
RUN apk --no-cache add ca-certificates tzdata \
			&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
			&& echo "Asia/Shanghai" >  /etc/timezone \
			&& apk del tzdata

WORKDIR /ignite-agent
COPY --from=builder /go/src/ignite-agent/cmd/ignite-agent/ignite-agent .

EXPOSE 4000
ENTRYPOINT ["./ignite-agent"]
