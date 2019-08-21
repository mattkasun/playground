FROM golang:latest as builder
WORKDIR /go/src/github.com/mattkasun/playground
RUN go get github.com/gin-gonic/gin
RUN go get github.com/dchest/uniuri
RUN go get golang.org/x/crypto/bcrypt

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix temp -ldflags '-extldflags "-static"' .


FROM busybox
WORKDIR /root/
COPY --from=builder /go/src/github.com/mattkasun/playground/playground .
ADD /resources/* resources/
ADD /stylesheet/css* stylesheet/css/
ADD /stylesheet/webfonts/* webfonts/
ADD /data/* data/
ADD /html/* html/
CMD ["./playground"]

