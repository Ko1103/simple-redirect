# Simple Redirect Container

## How to run

```shell
docker compose up
```

## Set to k8s

There are 2 things you need to aware of:

1. Set envoriment variable `REDIRECT_URL` in `k8s/deployment.yaml`

2. Set health check path in `k8s/deployment.yaml`

You can see example in `example/k8s/deployment.yaml`
