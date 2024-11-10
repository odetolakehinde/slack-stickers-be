FROM golang:1.23-alpine
RUN apk add git

ARG app_env
ENV APP_ENV $app_env

RUN touch /tmp/runner-build-errors.log
COPY . /go/src/github.com/odetolakehinde/slack-stickers-be
WORKDIR /go/src/github.com/odetolakehinde/slack-stickers-be/src

RUN go get ./
RUN go build
RUN go get github.com/pilu/fresh

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD if [ ${APP_ENV} = production ]; \
	then \
	src; \
	else \
	fresh; \
	fi

EXPOSE 6001
