FROM golang:1.23-alpine
RUN apk add git

ARG app_env
ENV APP_ENV $app_env

# Disable VCS stamping globally for Go
ENV GOFLAGS=-buildvcs=false

RUN touch /tmp/runner-build-errors.log
COPY . /go/src/github.com/odetolakehinde/slack-stickers-be
WORKDIR /go/src/github.com/odetolakehinde/slack-stickers-be/src

RUN go mod download
RUN go get ./
RUN go build -o /usr/local/bin/src  # Install the built binary to /usr/local/bin/
RUN go install github.com/zzwx/fresh@latest

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD if [ "${APP_ENV}" = "production" ]; then /usr/local/bin/src; else fresh; fi

EXPOSE 6001
