# Render kustomization template without applying
```shell
kubectl kustomize ./nginx/
```

# Apply rendered template
```shell
kubectl apply -k nginx
```
