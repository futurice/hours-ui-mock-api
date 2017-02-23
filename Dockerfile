FROM golang:1.6-alpine

RUN apk --no-cache add \
    git

RUN go get -d -v github.com/julienschmidt/httprouter

COPY . src/github.com/futurice/hours-ui-mock-api

ENV PORT 8000
EXPOSE 8000

RUN go install -v github.com/futurice/hours-ui-mock-api

ENTRYPOINT bin/hours-ui-mock-api
