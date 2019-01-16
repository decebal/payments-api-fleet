FROM golang:alpine

RUN apk update \
    && apk add --virtual build-dependencies \
        build-base \
        gcc \
        wget \
        git

ENV GO111MODULE=on

ENV APP_NAME go-docker
ENV REPO_BASE_URL github.com/decebal
ENV APP_PATH $GOPATH/src/$REPO_BASE_URL/$APP_NAME/payments-api-fleet

COPY . $APP_PATH
WORKDIR $APP_PATH
RUN make build-deps

CMD ["sh", "-c", "go run $(echo ${APP_PATH}/src/app.go)"]
