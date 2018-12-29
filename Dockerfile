FROM golang:latest as builder
WORKDIR /go/src/github.com/mattkasun/playground
RUN go get github.com/gin-gonic/gin
RUN go get github.com/dchest/uniuri

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix temp -ldflags '-extldflags "-static"' .


FROM busybox
WORKDIR /root/
COPY --from=builder /go/src/github.com/mattkasun/playground/playground .
ADD /resources/* resources/
ADD /stylesheet/* stylesheet/
ADD /html/*gohtml html/
CMD ["./playground"]

