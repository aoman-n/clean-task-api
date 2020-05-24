# common build
FROM golang:1.13 as base
COPY ./api /api
WORKDIR /api
RUN apt-get update && apt-get -y install default-mysql-client

# development build
FROM base as development

EXPOSE 8080
RUN go mod download
RUN GO111MODULE=off go get -u github.com/oxequa/realize && \
    go get -u bitbucket.org/liamstask/goose/cmd/goose && \
    go get -u github.com/kyoh86/richgo
ENTRYPOINT ["sh", "./entrypoint.sh"]
