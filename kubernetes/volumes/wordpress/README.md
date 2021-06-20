# Dynamic provisioning of storage 

## Create storage class
Choose either one, not both

#### Create storage class (for AWS)
```shell
kubectl apply -f storage-aws.yml
```

#### Create storage class (for local provisioner)
```shell
kubectl apply -f storage-local.yml
```

# Apply the manifests
```shell
kubectl apply -f storage-pvc.yml -f wordpress-db-service.yml -f wordpress-db.yml -f wordpress-secrets.yml -f wordpress-web-service.yml -f wordpress-web.yml
```
