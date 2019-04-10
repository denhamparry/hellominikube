# Hello Minikube - Go

A simple web application to use in a Minikube demo, details can be found [here](../README.md).

## Prerequisites

* [Go](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/install/)

## Build

### Go

To run the application locally with a terminal:

```bash
$  go build -o hellominikube .

```

### Docker

To build a Docker image:

```bash
$ docker build -t denhamparry/hellominikube:0.1 .
Sending build context to Docker daemon   7.68kB
Step 1/6 : FROM golang:latest
 ---> c7942203692b
Step 2/6 : RUN mkdir /app
 ---> Using cache
 ---> 203565ed0125
Step 3/6 : ADD . /app/
 ---> Using cache
 ---> 9d00bbe54609
Step 4/6 : WORKDIR /app
 ---> Using cache
 ---> 581db1e476b8
Step 5/6 : RUN go build -o main .
 ---> Using cache
 ---> 50541036cd78
Step 6/6 : CMD ["/app/main"]
 ---> Using cache
 ---> f77886c6500f
Successfully built f77886c6500f
Successfully tagged denhamparry/hellominikube:0.1
```

#### Docker Multi-stage

To help show the concept of a multistage build:

```bash
$ docker build -t denhamparry/hellominikube:0.2 -f Dockerfile.scratch .
Sending build context to Docker daemon  5.632kB
Step 1/14 : FROM golang@sha256:8cc1c0f534c0fef088f8fe09edc404f6ff4f729745b85deae5510bfd4c157fb2 as builder
 ---> d4953956cf1e
Step 2/14 : RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
 ---> Using cache
 ---> c870fcb3fcb0
Step 3/14 : RUN adduser -D -g '' appuser
 ---> Using cache
 ---> 4e121d11e3b4
Step 4/14 : WORKDIR /app
 ---> Using cache
 ---> 17e78ac9bb2a
Step 5/14 : COPY . .
 ---> 336c00e6fcff
Step 6/14 : RUN go get -d -v
 ---> Running in 5c04aa0a8cb4
Removing intermediate container 5c04aa0a8cb4
 ---> 4eaf59ecbfca
Step 7/14 : RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/hellominikube .
 ---> Running in 388eb6c32f49
Removing intermediate container 388eb6c32f49
 ---> 8164c529688e
Step 8/14 : FROM scratch
 --->
Step 9/14 : COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
 ---> Using cache
 ---> aea6c87b4ddc
Step 10/14 : COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
 ---> Using cache
 ---> fe701d4517e1
Step 11/14 : COPY --from=builder /etc/passwd /etc/passwd
 ---> Using cache
 ---> b2c767f8f17b
Step 12/14 : COPY --from=builder /go/bin/hellominikube /go/bin/hellominikube
 ---> 531af82d84b9
Step 13/14 : USER appuser
 ---> Running in f5b8a00dbfdb
Removing intermediate container f5b8a00dbfdb
 ---> 984cc1963f88
Step 14/14 : ENTRYPOINT ["/go/bin/hellominikube"]
 ---> Running in fb2e58ca255a
Removing intermediate container fb2e58ca255a
 ---> ec4dd5cfba96
Successfully built ec4dd5cfba96
Successfully tagged denhamparry/hellominikube:0.2
```

## Publish

### Docker

If you want to publish your Docker image:

```bash
$ docker push denhamparry/hellominikube:0.1
The push refers to repository [docker.io/denhamparry/hellominikube]
e34bc220e5f9: Pushed
f59f502a1f7e: Pushed
bc6cd85894a8: Pushed
bf8dd312214d: Mounted from library/golang
387f402927ce: Mounted from library/golang
f4907c4e3f89: Mounted from library/golang
b17cc31e431b: Mounted from library/golang
12cb127eee44: Mounted from library/golang
604829a174eb: Mounted from library/golang
fbb641a8b943: Mounted from library/golang
0.1: digest: sha256:1760bef6b2def6c270f2008df8835354c0a416f3674adbd591f0416e0a0c7f14 size: 2421
```

### Docker Multi-stage

If you want to pubish your Docker multi-stage build image:

```bash
$ docker push denhamparry/hellominikube:0.2
docker push denhamparry/hellominikube:0.2
The push refers to repository [docker.io/denhamparry/hellominikube]
e7177fa34705: Pushed
04bd29a3e5c9: Pushed
c0a354fa5e50: Pushed
b75a105926ff: Pushed
0.2: digest: sha256:20834076f520ec08d88e9a715933d66834cf5a91146d98902ed1c402e780d07d size: 1155
```

## Run

### Go

If you have run the previous Go build:

```bash
./hellominikube
```

Otherwise:

```bash
$ go run main.go
```

Then visit the URL [localhost:5000](http://localhost:5000) in a browser or cURL:

```bash
$  curl localhost:5000
Hello Minikube! (v0.1)
```

### Docker

To run the Docker image locally:

```bash
$ docker run -d -p 5000:5000 denhamparry/hellominikube:0.1

```

### Docker Multi-stage

To run the Docker multi-stage build image locally:

```bash
$ docker run -d -p 5000:5000 denhamparry/hellominikube:0.2

```