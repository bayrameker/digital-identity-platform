
installations

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update

helm install api-gateway ingress-nginx/ingress-nginx \
  --namespace ingress-nginx --create-namespace



kubectl create secret tls tls-secret \
  --key path/to/server.key \
  --cert path/to/server.crt \
  --namespace api-gateway



ek olarak konfigler : 

metadata:
  annotations:
    nginx.ingress.kubernetes.io/enable-rate-limiting: "${enable-rate-limiting}"
    nginx.ingress.kubernetes.io/limit-rps: "${rate-limit}"



Ingress Anotasyonları ile

metadata:
  annotations:
    nginx.ingress.kubernetes.io/whitelist-source-range: "192.168.1.0/24"
    nginx.ingress.kubernetes.io/limit-rps: "50"
    nginx.ingress.kubernetes.io/limit-burst-multiplier: "3"


Versiyonlandırılmış Servisler:
user-service-v1
user-service-v2


rules:
  - host: api.example.com
    http:
      paths:
        - path: /api/v1/users(/|$)(.*)
          pathType: Prefix
          backend:
            service:
              name: user-service-v2
              port:
                number: 3000



Argo CD Kurulumu:

kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml



Prometheus Operator Kurulumu:

kubectl create namespace monitoring
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install prometheus prometheus-community/kube-prometheus-stack --namespace monitoring
