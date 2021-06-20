# Creating a base for a development and production Deployment

## Render overlays without applying
```shell
kubectl kustomize ./overlays/dev
kubectl kustomize ./overlays/prod
```

## Apply manifest 
```shell
kubectl apply -k ./overlays/dev
kubectl apply -k ./overlays/prod
```