FROM golang:alpine

ENV APP_NAME go-docker
ENV REPO_BASE_URL github.com/decebal
ENV APP_PATH $GOPATH/src/$REPO_BASE_URL/$APP_NAME/payments-api-fleet

COPY . $APP_PATH
WORKDIR $APP_PATH
RUN go get -u

CMD ["sh", "-c", "go run $(echo ${APP_PATH}/src/app.go)"]
