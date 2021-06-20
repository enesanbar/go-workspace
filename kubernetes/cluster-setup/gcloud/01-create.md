## Install command line tools
```shell
# download the sdk
curl https://sdk.cloud.google.com | bash

# follow the initialization prompts.
gcloud init

# enable compute engine api
gcloud services enable compute.googleapis.com
```

## Create managed kubernetes cluster
```shell
gcloud container clusters create test-cluster-1 \
--cluster-version latest --machine-type n1-standard-2 \
--image-type UBUNTU --disk-type pd-standard --disk-size 100 \
--no-enable-basic-auth --metadata disable-legacy-endpoints=true \
--scopes compute-rw,storage-ro,service-management,service-control,logging-write,monitoring \
--num-nodes "3" --enable-stackdriver-kubernetes \
--no-enable-ip-alias --enable-autoscaling --min-nodes 1 \
--max-nodes 5 --enable-network-policy \
--addons HorizontalPodAutoscaling,HttpLoadBalancing \
--enable-autoupgrade --enable-autorepair --maintenance-window "10:00"
```

* --cluster-version sets the Kubernetes version (default: latest) 

Get the available k8s versions
```shell
gcloud container get-server-config
```

* --machine-type parameter sets the instance type (default: n1-standard-1).

Get the list of predefined types
```shell
gcloud compute machine-types list
```

* --image-type sets OS image (default: COS)

Get the list of available image types
```shell
gcloud container get-server-config
```

* --enable-autoupgrade enables the GKE auto-upgrade feature for cluster nodes and 
* --enable-autorepair enables the automatic repair feature, which is started at the time defined with the --maintenance-window parameter. 

# Deploying with a custom network configuration
* Create a VPC network
```shell
gcloud compute networks create test-kubernetes-network --subnet-mode custom
```

* Create a subnet in your VPC network
```shell
gcloud compute networks subnets create test-kubernetes-subnetwork --network test-kubernetes-network --range 10.240.0.0/16
```

* Create a firewall rule to allow internal traffic
```shell
gcloud compute firewall-rules create test-kubernetes-allow-int \
--allow tcp,udp,icmp --network test-kubernetes-network \
--source-ranges 10.240.0.0/16,10.200.0.0/16
```

* Create a firewall rule to allow external SSH, ICMP, and HTTPS traffic:
```shell
gcloud compute firewall-rules create test-kubernetes-allow-ext \
--allow tcp:22,tcp:6443,icmp --network test-kubernetes-network \
--source-ranges 0.0.0.0/0
```

* Verify rules
```shell
gcloud compute firewall-rules list
```

* Deploy as before with following parameters
```shell
gcloud container clusters create test-cluster-1
...
--network test-kubernetes-network \ 
--subnetwork test-kubernetes-subnetwork
```

# Deleting cluster
* Get a list of clusters to get the name of the cluster
```shell
gcloud container clusters list
```

* Delete the cluster with name
```shell
gcloud container clusters delete CLUSTER-NAME
```
