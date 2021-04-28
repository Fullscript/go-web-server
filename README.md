# go-web-server

## Requirements

Write a Dockerfile that compiles and then runs the web server in this repository.  

Please write it as if the container will be used in production.  

1. The Dockerfile must build an image for this go-web-server successfully.
2. The image must run successfully and have the go-web-server listening on the configured port.
3. Please provide the response returned from the `/keyword` route.

## Details

It exposes a web server on port 3030 and logs to STDOUT.  The port is configurable by setting the PORT environment variable.  

It has several routes that return status code 200 and some data: `/health`, `/hello`, and `/keyword`. All other routes will return 404: "404 page not found".  
