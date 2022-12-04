# apigw-nginx

A bunch of examples with custom docker images, docker files and nginx configurations

### whoami microservice

It's a dummy microservice written in Go. It exposes some endpoints as HTTP Server. In order to emulate different microservice, it 
exposes `\whoami` endpoint that return the (container) hostname. In docker-compose with 1 replica per microservice, it is possible to set the container_name and the hostname (that, by default doesn't match)

If you will run load-balancing examples, you'll not be able to set container_name (mismatching with replicas) and the hostname will match with `container_id` variable


## Getting started

```
$ cd docker
$ docker-compose -f docker-compose.$EXAMPLE_NAME.yml up -d
$ curl -iL http://localhost:8080/orders/whoami
```
