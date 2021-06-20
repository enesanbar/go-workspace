## Install and configure awscli
```shell
brew install awscli
aws configure
```

## Install eksctl
```shell
brew tap weaveworks/tap
brew install weaveworks/tap/eksctl
```

## Create cluster
```shell
# managed master
eksctl create cluster --name=cluster-1 --nodes=2 --region=eu-west-1

# managed master + nodes
eksctl create cluster --name=cluster-1 --nodes=2 --region=eu-west-1 --managed
```

eksctl should update your kube config and set your current context to the cluster created.

## Confirm the cluster version and workers
```shell
kubectl cluster-info
kubectl get nodes -o wide
```

# Setup IAM Roles for Service Accounts

Enable IAM Roles for Service Accounts on the EKS cluster

```
eksctl utils associate-iam-oidc-provider --cluster=cluster-2
eksctl utils associate-iam-oidc-provider --cluster=cluster-2 --approve
```

Create new IAM Role using eksctl
```
eksctl create iamserviceaccount --cluster=cluster-2 --name=myserviceaccount --namespace=default --attach-policy-arn=arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess
eksctl create iamserviceaccount --cluster=cluster-2 --name=myserviceaccount --namespace=default --attach-policy-arn=arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess --approve
```
