# simple_http

## Tutorial

start server by docker:
```shell
docker pull willerhe/simple-http:latest
docker run -d -p 80:80 willerhe/simple-http:latest
```

usecase:
```shell
curl localhost:80/healthz
curl localhost:80 -H "TestHeader: willerhe, abc"
```