# EKS K8S README

### Setup EKS Fargate cluster.

```
eksctl create cluster --config-file ./cluster-fargate.yaml

CLUSTER_NAME=fortune-cluster
```

### Setup AWS Load Balancer Controller

Refs:

* https://docs.aws.amazon.com/eks/latest/userguide/aws-load-balancer-controller.html
* https://www.eksworkshop.com/beginner/180_fargate/prerequisites-for-alb/

```
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
  -n kube-system

kubectl get deployment -n kube-system aws-load-balancer-controller

```