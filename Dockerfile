FROM golang:1.15-alpine
LABEL maintainer="Juan Manuel <jumacasni@correo.ugr.es>"

RUN apk --no-cache add curl \
	&& apk add --no-cache bash \
	&& apk add --no-cache build-base \
	&& curl -sL https://git.io/tusk | bash -s -- -b /usr/local/bin latest

COPY internal/ tusk.yml /terrake/

WORKDIR /terrake

CMD ["tusk", "test"]