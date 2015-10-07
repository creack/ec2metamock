FROM       golang:1.5
MAINTAINER Guillaume J. Charmes <guillaume@charmes.net>

ENV        APP_DIR    $GOPATH/src/github.com/creack/ec2metamock
WORKDIR    $APP_DIR

ADD        . $APP_DIR
RUN        go build

ENTRYPOINT ["./ec2metamock"]
