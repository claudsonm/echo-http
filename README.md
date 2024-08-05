# Echo HTTP

Container image that spins up a web server and echoes in the console and screen the request data.

# How to build and execute

```bash
# Building
docker buildx build --platform linux/amd64 --load -t claudsonm/echo-http .
docker push claudsonm/echo-http:latest

# Running
docker run -p 8080:80 -d --rm -t claudsonm/echo-http
```
