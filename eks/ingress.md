Create Nginx Controller Ingress
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.11.1/deploy/static/provider/aws/deploy.yaml
```

See more at: https://kubernetes.github.io/ingress-nginx/deploy/#aws

Check nameserver of a domain
```
host -t ns leosite.cloud
```