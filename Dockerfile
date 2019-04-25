FROM golang:alpine

RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh

WORKDIR /app

COPY . ./

RUN ["bash", "./install.sh"]

EXPOSE 8080

CMD [ "go", "run", "main.go" ]
