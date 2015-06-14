# Fortune Cookie API

A simple API to return fortunes. Written in Go and deployed to AppEngine Managed VMs. The deployment target (AppEngine Managed VMs) necessitates some changes within the Go app from the standard library and using AppEngine, in general, necessitates familiarity with `app.yaml`.

AppEngine Managed VM golang app with project directory structure as per [gb](http://getgb.io/), an alternative Go build tool.

### Endpoints

* web: [gdgnoco-fortune.appspot.com](https://gdgnoco-fortune.appspot.com)
* API: [gdgnoco-fortune.appspot.com/api/fortune](https://gdgnoco-fortune.appspot.com/api/fortune)


## Prerequisites

* Docker
* `gcloud`
* Go

```
gcloud components update gae-go app
go get -u google.golang.org/appengine
```

## Build

A Google Cloud Project will be necessary at this point to do a trial build or development.

Google Cloud Project `gdgnoco-fortune` (refered to as `PROJECT_NAME`)

First login with `gcloud config set project` then run the local Docker vm via `gcloud preview app run`

```
gcloud config set project PROJECT_NAME
gcloud preview app run ./app.yaml
```

## Deploy

Deploying is done via the `gcloud preview app deploy` command.

```
gcloud preview app deploy ./app.yaml
```



## To Do

* parse out `\n\t\t-- AUTHOR\n` as author, add to Fortune struct
* nice front web page, use React/Riot/Polymer/Angular to call api and show fortune
* unescape braindead string parsing
* use other [features of Google AppEngine](https://cloud.google.com/appengine/features/) ([Datastore](https://cloud.google.com/appengine/docs/go/gettingstarted/usingdatastore)) to allow custom fortunes

## Notes

FYI, what makes an AppEngine app a Managed VM app are the following additions to `app.yaml`

```
runtime: go
vm: true
api_version: go1
```

### Dockerfile

Basic gcr Dockerfile looks in the project root folder. Change the `COPY` command to point to the local package source `src/gdg-fortunecookieapi` (since we're using a [gb project structure](http://getgb.io/))

```
COPY ./src/gdg-fortunecookieapi /go/src/app
```

### Docker maintenance
To remove exited docker containers (on OSX w/ boot2docker, therefore no `sudo`)

```
docker rm $(docker ps -a | grep Exited | awk '{print $NF}')
```

To remove unassociated/untagged docker images

```
docker rmi $(docker images | grep "^<none>" | awk '{print $3}')
```




