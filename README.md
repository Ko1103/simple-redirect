# Simple Redirect Container

Simple container for redirect all request but health check path to specific url.

## Get started

You can run by docker-compose.

```shell
docker compose up -d
```

and access to `http://localhost:8080` you will be redirected to `https://www.google.com`

When you access to `http://localhost:8080/health` you will get `200 OK` response.

## Run in k8s

You can use this as pod in k8s.

full example is [here](https://github.com/Ko1103/simple-redirect/blob/main/example/k8s/deploy.yaml)

3 steps to run in k8s:

1. Set container: `michealko/simple-redirect:latest`
2. Set envoriment variable `REDIRECT_URL`. e.g. `https://www.google.com`.
3. Set health check path. e.g. `/health`

# Troubleshooting

If ingress health check is failed, it may be caused by service health check path. Please check the health check path in your ingress configuration.

For example, if you use GCP, use BackendConfig to set health check path.

```yaml
apiVersion: cloud.google.com/v1
kind: BackendConfig
metadata:
  name: simple-redirect-backendconfig
spec:
  healthCheck:
    requestPath: /health
```
