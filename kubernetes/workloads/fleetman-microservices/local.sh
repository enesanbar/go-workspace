kubectl apply -f api-gateway-deployment.yml -f api-gateway-svc.yml
kubectl apply -f storage-local.yml
kubectl apply -f mongo-deployment.yml -f mongo-svc.yml
kubectl apply -f position-simulator-deployment.yml
kubectl apply -f position-tracker-deployment.yml -f position-tracker-svc.yml
kubectl apply -f queue-deployment.yml -f queue-svc.yml
kubectl apply -f webapp-deployment.yml -f webapp-svc.yml
