FROM golang:latest as builder
WORKDIR /go/src/github.com/mattkasun/playground
COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .


FROM busybox

COPY --from=builder /go/src/github.com/playground/playground /
ADD /html/*gohtml /html/
ADD /data/*data /data/
CMD ["/playground"]

