FROM golang

# if left blank app will run with dev settings
# to build production image run:
# $ docker build ./server --build-args app_env=production
ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/github.com/conorburke/Seker/server
WORKDIR /go/src/github.com/conorburke/Seker/server

# added vendor services will need to be included here
RUN go get ./vendor/database
RUN go get ./vendor/mytest

RUN go get ./
RUN go build

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD if [ ${APP_ENV} = production ]; \
	then \
	server; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8080