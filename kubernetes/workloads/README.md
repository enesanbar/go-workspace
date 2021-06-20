# Scaffolding with generators
Use `--dry-run=client -o yaml` to generate yaml as a starting point

## Generate Deployment
```shell
kubectl create deployment nginx --image nginx --replicas 1 --dry-run=client -o yaml > deployment.yml
```

## Generate Service
```shell
kubectl create service clusterip nginx --tcp 80:80 --dry-run=client -o yaml > service.yml
```

See also: https://kubernetes.io/docs/reference/kubectl/conventions/#generators

