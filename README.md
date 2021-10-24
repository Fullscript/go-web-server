# go-web-server

## Requirements

Write a Dockerfile that compiles and then runs the web server in this repository.  

Please write it as if the container will be used in production.  

1. The Dockerfile must build an image for this go-web-server successfully.
2. The image must run successfully and have the go-web-server listening on port 3030.
3. Once you have the Dockerfile running **please provide the response** returned from the `/keyword` route to us.  

## Details

It exposes a web server on port 3030 and logs to STDOUT.  The port is configurable by setting the PORT environment variable.  

It has several routes that return status code 200 and some data: `/health`, `/hello`, and `/keyword`. All other routes will return 404: "404 page not found".  

## Developer Notes
There are 2 docker files within this repository:
1) `Dockerfile` - can be used to run a development environment on port 3030 - which will be linked to a local filestore and watch for any changes.
    *) To build this image: `docker build --build-arg USER_ID=1000 --build-arg GROUP_ID=1000 -t fullscript/go .`
        *) Where USER_ID and GROUP_ID are to ensure that the containers in the process are not running as root (if on linux can replace 1000 with `$(id- u)` and `$(id -g)` respectively).
    *) Once built to run this image: 
        *) Linux: `docker run -it --rm -p 3030:3030 -v $(pwd):/go/src fullscript/go`  
        *) PowerShell : `docker run -it --rm -p 3030:3030 -v ${PWD}:/go/src fullscript/go`
2) `Dockerfile.production` - can used to run a production environment on port 3030
    *) To build: `docker build -t fullscript/go-production -f .\Dockerfile.production .`
    *) To run: `docker run -it -p 3030:3030 fullscript/go-production`

## Keyword
The reponse from `/keyword` is `wellness`.