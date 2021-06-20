## Install and configure awscli
```shell
brew install awscli
aws configure
```

## Install kops
```shell
brew install kops
```

## Managing K8S Cluster with kops

### Create an S3 bucket for the kops state
```shell
BUCKET_NAME=eanbar-k8s-kops-state
aws s3api create-bucket --bucket ${BUCKET_NAME} --region us-east-1
aws s3api put-bucket-versioning --bucket ${BUCKET_NAME} --region us-east-1 --versioning-configuration Status=Enabled

export KOPS_STATE_STORE=s3://${BUCKET_NAME}
```

### Create a new cluster
##### Write config to store
> `kops create cluster` command doesn't create the clusters itself. It creates the configuration file in the S3 bucket.

```shell
# names ending with k8s.local creates a gossip-based cluster instead of relying on dns
export KOPS_CLUSTER_NAME=eanbar.k8s.local
kops create cluster --name ${KOPS_CLUSTER_NAME} \
    --node-size t2.micro \
    --node-count=3 \
    --zones eu-west-3a,eu-west-3b,eu-west-3c \
    --master-size t2.micro \
    --master-zones eu-west-3a,eu-west-3b,eu-west-3c

# or generate yaml config to apply it later with
# kops create -f 01-kops-config.yaml
kops create cluster --name ${KOPS_CLUSTER_NAME} \ 
    --node-size t2.micro \
    --node-count=3 \
    --zones eu-west-3a,eu-west-3b,eu-west-3c \
    --master-size t2.micro \
    --master-zones eu-west-3a,eu-west-3b,eu-west-3c \
    --dry-run -o yaml > 01-kops-config.yaml
```

##### Deploy generated config
```shell
kops create secret --name ${KOPS_CLUSTER_NAME} sshpublickey admin -i ~/.ssh/id_rsa.pub
kops update cluster --name ${KOPS_CLUSTER_NAME} --yes
```

##### Wait for cluster to be ready
```shell
kops validate cluster --wait 10m
```

Note: If you get "Unauthorized" errors, kubeconfig may not be properly updated. 
Update kubeconfig manually with the following command
```shell
kops export kubecfg --admin
```

### Update the cluster

#### Upgrading k8s to a specific version
Set kubernetesVersion key to the desired k8s version.

```shell
kops edit cluster --name ${KOPS_CLUSTER_NAME}
kops update cluster --name ${KOPS_CLUSTER_NAME} --yes
kops rolling-update cluster --name ${KOPS_CLUSTER_NAME} --yes
```

#### Upgrading k8s to the latest version
```shell
kops upgrade cluster --name ${KOPS_CLUSTER_NAME} --yes
kops update cluster --name ${KOPS_CLUSTER_NAME} --yes
kops rolling-update cluster --name ${KOPS_CLUSTER_NAME} --yes
```

#### Scaling nodes
Update instance groups configuration to set minSize and maxSize values.
```shell
kops edit instancegroups nodes --name ${KOPS_CLUSTER_NAME} --yes
kops edit instancegroups master-eu-west-3a --name ${KOPS_CLUSTER_NAME} --yes
kops edit instancegroups master-eu-west-3b --name ${KOPS_CLUSTER_NAME} --yes
kops update cluster --name ${KOPS_CLUSTER_NAME} --yes
kops rolling-update cluster --name ${KOPS_CLUSTER_NAME} --yes
```

### Delete the cluster

First delete the cluster
```shell
kops delete cluster --name ${KOPS_CLUSTER_NAME}  --yes
```

Then manually delete the S3 bucket
```shell

```