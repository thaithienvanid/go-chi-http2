FROM golang:alpine

RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh

WORKDIR /app

COPY . ./

RUN ["go", "build"]

EXPOSE 8080

CMD [ "./main" ]
