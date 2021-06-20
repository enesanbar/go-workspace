## Create deployment with volume attached
1. Manually create a volume in AWS
```shell
aws ec2 create-volume --region eu-west-3 --availability-zone eu-west-3a --volume-type gp2 --size 10 --tag-specifications 'ResourceType=volume, Tags=[{Key=KubernetesCluster,Vaue=kubernetes.test}]'
```

2. Use volume id in the deployment yaml
```shell
kubectl apply -f helloworld-deployment.yml
```
## Clean up

```shell
kubectl delete -f helloworld-deployment.yml
aws ec2 delete-volume --volume-id VOLUME_ID_HERE
```
