Using kubectl to connect to EKS Cluster
```
$ aws eks update-kubeconfig --name banking-go --region ap-southeast-1
>> Added new context arn:aws:eks:ap-southeast-1:730335272010:cluster/banking-go to /home/kha/.kube/config
```
Ensure that the current aws user configured on local has been in EKSFullAccess permission group

## Checking the current context of kubectl

```
$ cat ~/.kube/config
kubectl config use-context <<current-context>>
```
?? How to view the current aws configuration, use the following command:
$ cat ~/.aws/credentials 
# Result would be like:
[default]
aws_access_key_id = <accessKey>
aws_secret_access_key = <secretKey>
```

## Checking the kubectl client version
Make sure that kubectl client belongs to EKS
```
$ kubectl version --client
Client Version: v1.30.0-eks-036c24b
Kustomize Version: v5.0.4-0.20230601165947-6ce0bf390ce3
```

So now we can work with AWS EKS through kubectl on the local machine
```
$ kubectl cluster-info
Kubernetes control plane is running at https://A991BB83B6CC6F9F18EABB1F3AFC1843.gr7.ap-southeast-1.eks.amazonaws.com
CoreDNS is running at https://A991BB83B6CC6F9F18EABB1F3AFC1843.gr7.ap-southeast-1.eks.amazonaws.com/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
```

See more here: https://repost.aws/knowledge-center/amazon-eks-cluster-access\


## More infomation about the current aws profile on local machine
Here is the current credentials on local machine:
```
$ cat ~/.aws/credentials 
[default]
aws_access_key_id = <root_access_key>
aws_secret_access_key = <root_secret_key>
[github]
aws_access_key_id = <iam_user_access_key>
aws_secret_access_key = <iam_user_secret_key>
```

And if we wanna change the current profile on client, using the env $AWS_PROFILE
```
export AWS_PROFILE=default
```


## Implement TTL with Kubectl
```
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.15.1/cert-manager.yaml
```