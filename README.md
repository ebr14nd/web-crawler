# web-crawler

Hey :wave:

Here is my very simple and naive web-crawler!

# Start server

Directly with golang cli:
```shell
go run cmd/wcserver/main.go
```
With docker cli
```shell
docker container run ebr41nd/wc-server
```

# Start cli

By default, the client calls the server at the `http://localhost:80` address. It can be changed by setting the WC_SERVER_HOST environment variable

Directly with golang cli:
```shell
go run cmd/wcclient/main.go https://example.com
```
With docker cli
```shell
docker container run ebr41nd/wc-client https://example.com
```

# Build

In the Makefile, there is the `build-wcserver` and `build-wcclient` to build the container images of server and client 

# Deployment

In the Makefile, to deploy on Kubernetes, use the `deploy-wcserver` and `deploy-wcclient` targets

The server is deployed as a Deployment, a `wc-server` Service is created to be able to access the server.

The client is deployed as a Job. Since there is only one env variable, I didn't used a ConfigMap and directly set an environment variable in the Job's PodTemplate

# Assumptions

I made the follow assumptions: 
- when an URL is given, only children paths are followed :
```
In http://example.com page, there is data link, http://example.com/data is crawled
```
```
In http://example.com/data, there is http://example/more link, it is not crawled
```
```
In http://example.com/, there is https://redhat.com link, it is not crawled
```
- If there is a circular reference, the web-crawler loop :sweat:
- I made fews issues of enhancements since I won't be able to implement next steps :smile:

# Code organisation

- `/cmd` contains the main packages of client / server
- `/internal` contains all the code of client / server
- `/k8s` contains the Kubernetes descriptors
- `/testdata` contains test files


