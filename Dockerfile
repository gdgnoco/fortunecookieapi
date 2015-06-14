# Dockerfile extending the generic Go image with application files for a
# single application.
FROM gcr.io/google_appengine/golang

RUN apt-get update && apt-get install -y fortunes

COPY ./src/gdg-fortunecookieapi /go/src/app
COPY ./web /go/src/app/web

RUN go get -u google.golang.org/appengine
RUN go-wrapper download
RUN go-wrapper install -tags appenginevm
