# 查看所有node的label
kubectl get nodes -o wide

kubectl get nodes c1-worker -o jsonpath='{.metadata.labels}'
kubectl get nodes c1-worker2 -o jsonpath='{.metadata.labels}'
kubectl get nodes c1-worker3 -o jsonpath='{.metadata.labels}'

"kubernetes.io/hostname":"c1-worker"

# 将hello world 部署到对应的node上

kubectl apply -f appyaml/httpbin.yaml


# 测试api

istio/helloworld
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"


# 测试步骤
1. 准备kind cluster
2. `bash labellocality.sh`给node标记locality
3. 安装kmesh
3. 在该目录下 `bash test.sh`
