FROM golang:1.9.0
WORKDIR /go/src/murrman/audio-share
RUN go get github.com/labstack/echo
RUN go get github.com/labstack/echo/middleware
COPY . /go/src/murrman/audio-share
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /home/
COPY --from=0 /go/src/murrman/audio-share .
CMD [ "./main" ]