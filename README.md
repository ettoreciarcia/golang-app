# Golang-app

A sample app to test Kubernetes Resource

A complete build command:

```docker build -t [IMAGE_NAME] . --build-arg VERSION_NUMBER=1 --build-arg PORT_NUMBER=8080```

This container can handle HTTP request on path ```/```
Example of response

```json
{
    "hostname": "170e800c2068",
    "message": "Status Created",
    "version": "1"
}
```

You can find the image at [hecha00/golang-app](https://hub.docker.com/repository/docker/hecha00/golang-app)


Do you want try it? 

```docker run -dit -p 8080:8080 hecha00/golang-app:latest```

Yeah. i know. I shouldn't be using the latest tag but no production environment will be on fire, I promise

I will use this container to test Kubernetes Services :)