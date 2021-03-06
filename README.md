Minimal Go application container (¿unikernel? ¿nano/fempto-service? 
===================================================================

How to build the minimal docker container for a go application. Thos container is based from **scratch**. 

When running, the Go binary is looking for libraries on the operating system it’s running in. We compiled our app, but it still is dynamically linked to the libraries it needs to run (i.e., all the C libraries it binds to). Unfortunately, scratch is empty, so there are no libraries and no loadpath for it to look in. What we have to do is modify the build script to statically compile the app with all libraries built in. 

With this two little things (*from scratch container* and *statically linked application*) the result is an about **6.563 MB** container!!

Based on article [Building Minimal Docker Containers for Go Applications](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/).

This application is a minimal RESTful HTTP/2 service.


## Platform requirements

  - Go (tested with 1.5.1)
  - Docker (tested with 1.9.0)


## How to build the application 

Usually to build the program:

    go build -o rest-server


But this make a artifact with shared references. To build a static linked executable:

    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rest-server .


## How to build the container

    docker build -t nano-container .


## How to run the container

    docker run --rm -ti -p 8443:8443 nano-container


## Running the application

The application has a HTTPS entry point (https://localhost:8443/) that returns a simple JSON message:

    {
        "platform": "docker",
        "language": "go",
        "result": "bazinga!"
    }

The server console shows thr request remote address and protocol

Example:

    https://localhost:8443

    curl --insecure https://localhost:8443


# To HTTP/2

This example uses a auto-generated SSL certificate.

	http://intogooglego.blogspot.com.es/2015/09/day-20-go-http2-server-example.html
	www.ianlewis.org/en/http2-and-go

Download http/2 library
    
    go get golang.org/x/net/http2





