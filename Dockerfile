FROM golang:1.16 as builder
 
WORKDIR /go/src/github.com/wang-sy/cluster-manager
COPY . .
 
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.io && \
    go build -tags netgo -o service ./main.go
 
FROM busybox
 
COPY --from=builder /go/src/github.com/wang-sy/cluster-manager/service /service
 
EXPOSE 80
ENTRYPOINT [ "/service" ]
