# Download azure cli
```shell
brew install azure-cli
az --version
```

# Provision a managed cluster

* Log in to your account
```shell
az login
```

* Create a resource group named k8s-testing
```shell
az group create --name k8s-testing --location eastus
```

* Create a service principal and take note of your appId and password for the next steps
```shell
az ad sp create-for-rbac --skip-assignment
```

* List available k8s versions
```shell
az aks get-versions --location eastus --output table
```

* Create a cluster
```shell
az aks create --resource-group k8s-testing \
 --name AKSCluster \
 --kubernetes-version 1.15.4 \
 --node-vm-size Standard_DS2_v2 \
 --node-count 3 \
 --service-principal <appId> \
 --client-secret <password> \
 --generate-ssh-keys
```

* Connecting to AKS cluster
```shell
az aks get-credentials \ 
  --resource-group k8s-testing \
  --name AKSCluster
```

* Deleting the cluster
```shell
az aks delete --resource-group k8s-testing --name AKSCluster
```

## Notes
* set default group to specify resource group in every command
```shell
az configure --defaults group=k8s-tesing
```

