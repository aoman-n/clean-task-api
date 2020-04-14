# common build
FROM golang:1.13 as base
COPY . /api
WORKDIR /api
RUN apt-get update && apt-get -y install default-mysql-client

# development build
FROM base as development
ENV GO_ENV=development
EXPOSE 8080
RUN go mod download \
    && go get github.com/oxequa/realize
ENTRYPOINT ["sh", "./entrypoint.sh"]
