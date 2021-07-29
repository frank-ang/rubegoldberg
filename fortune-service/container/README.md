# EKS K8S README

## Minikube on localhost

### Minikube Hello

```bash
minikube start
minikube dashboard --url
kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4
kubectl get deployments
kubectl get pods
kubectl get events
kubectl config view
kubectl expose deployment hello-node --type=LoadBalancer --port=8080
kubectl get services
# Clean up
kubectl delete service hello-node
kubectl delete deployment hello-node
minikube stop
```

### Minikube Fortune.

Basic healthcheck.

```bash
# Setup
eval $(minikube docker-env)
minikube start

# 1-time setup, enable local minikube registry
minikube addons enable registry

# Setup Environment, build to registry.
make build-docker

# Edit deployment yaml imagePullPolicy 
## Set from Always to Never (the image is assumed to exist locally).
## imagePullPolicy: Never

# Deploy
kubectl apply -f fortune-single.yaml
kubectl get deployment fortune

# Test opening the service endpoint, it should show the healthcheck message.
minikube service fortune
```

Next, run with environment variables.

```bash
make set-kube-config 
make build-kube 
make  test-kube
```

## EKS

### Setup EKS Fargate cluster.

```bash
eksctl create cluster --config-file ./cluster-fargate.yaml
AWS_REGION=ap-southeast-1
VPC_ID=vpc-0aab98e7906fc81e5
CLUSTER_NAME=fortune-cluster
```

### Setup AWS Load Balancer Controller

Refs:

* https://docs.aws.amazon.com/eks/latest/userguide/aws-load-balancer-controller.html
* https://www.eksworkshop.com/beginner/180_fargate/prerequisites-for-alb/

```bash

curl -o iam_policy.json https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v2.2.0/docs/install/iam_policy.json

aws iam create-policy \
  --policy-name AWSLoadBalancerControllerIAMPolicy \
  --policy-document file://iam_policy.json

POLICY_ARN=SET_ME e.g. arn:aws:iam::111122223333:policy/AWSLoadBalancerControllerIAMPolicy

eksctl utils associate-iam-oidc-provider --region=ap-southeast-1 --cluster=$CLUSTER_NAME --approve

eksctl create iamserviceaccount \
  --cluster=$CLUSTER_NAME \
  --namespace=kube-system \
  --name=aws-load-balancer-controller \
  --attach-policy-arn=$POLICY_ARN \
  --override-existing-serviceaccounts \
  --approve

kubectl apply -k "github.com/aws/eks-charts/stable/aws-load-balancer-controller/crds?ref=master"

helm repo add eks https://aws.github.io/eks-charts

helm repo update

helm upgrade -i aws-load-balancer-controller eks/aws-load-balancer-controller \
  --set clusterName=$CLUSTER_NAME \
  --set serviceAccount.create=false \
  --set serviceAccount.name=aws-load-balancer-controller \
  -n kube-system \
  --set region=${AWS_REGION} \
  --set vpcId=${VPC_ID}

kubectl get deployment -n kube-system aws-load-balancer-controller
kubectl -n kube-system rollout status deployment aws-load-balancer-controller
```


Update Local CLI
```bash
# update local CLI
aws eks update-kubeconfig --name fortune-cluster

# uncurated commands, just taking note for now...
aws eks get-token --cluster-name fortune-cluster | jq -r '.status.token'


```

### Enable SSO user for kubectl

As the SSO user, run ```aws sts get-caller-identity```, and locate the corresponding role in IAM console. E.g. "AWSReservedSSO_...", removing the leading  “/sso.amazonaws.com/<region>” 

As the creator of the cluster, add the SSO role into the config map
```bash
kubectl edit configmap aws-auth -n kube-system
```

Append the role:

```yaml
  mapRoles: |
    ...<append below>...
    - rolearn: arn:aws:iam::<accountID>:role/AWSReservedSSO_AdministratorAccess_REDACTED
      username: admin:{{SessionName}}
      groups:
        - system:masters
```

Now, verify the SSO user should be able to execute ```kubectl``` commands successfully, and also view cluster details in the EKS Console.

```bash
kubectl get svc
kubectl get nodes -o wide
kubectl get pods --all-namespaces -o wide
```

## Deploy Fortune service to EKS

Fargate Profile setup.

```bash
eksctl create fargateprofile --cluster fortune-cluster \
    --name fortune  --namespace fortune
eksctl get fargateprofile --cluster fortune-cluster -o yaml
```

Did not change CoreDNS yet...