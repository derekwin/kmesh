bash labellocality.sh

kubectl apply -f sleep-c1-w3.yaml
kubectl apply -f helloworld-srv.yaml
kubectl apply -f helloworld-region1.zone1-1.yaml
kubectl apply -f helloworld-region1.zone1-2.yaml
kubectl apply -f helloworld-region1.zone2-1.yaml
kubectl apply -f helloworld-region1.zone2-2.yaml
kubectl apply -f helloworld-region2.zone3-1.yaml

sleep 6
echo ""
echo "################################################"
echo "在sleep-c1-w3（对应region2 zone3）访问api，测试五次:"

kubectl get pods -o wide

kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"

echo ""
echo "################################################"
echo "向region2 zone3新增一个后端，sleep-c1-w3访问api，测试五次:"

kubectl apply -f helloworld-region2.zone3-2.yaml
sleep 5
kubectl get pods -o wide

kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"

echo ""
echo "################################################"
echo "删掉region2 zone3所有后端，sleep-c1-w3访问api，测试十次:"s

kubectl delete -f helloworld-region2.zone3-2.yaml
kubectl delete -f helloworld-region2.zone3-1.yaml
sleep 5
kubectl get pods -o wide

kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"

echo ""
echo "################################################"
echo "恢复region2 zone3所有后端，sleep-c1-w2访问api，测试五次:"

kubectl apply -f helloworld-region2.zone3-2.yaml
kubectl apply -f helloworld-region2.zone3-1.yaml
kubectl delete -f sleep-c1-w3.yaml
kubectl apply -f sleep-c1-w2.yaml
sleep 6
kubectl get pods -o wide

kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"

echo ""
echo "################################################"
echo "删掉region1 zone2所有后端，sleep-c1-w2访问api，测试五次:"

kubectl delete -f helloworld-region1.zone2-2.yaml
kubectl delete -f helloworld-region1.zone2-1.yaml
sleep 5
kubectl get pods -o wide

kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"

echo ""
echo "################################################"
echo "继续删掉region1 zone1所有后端，sleep-c1-w2访问api，测试五次:"

kubectl delete -f helloworld-region1.zone1-2.yaml
kubectl delete -f helloworld-region1.zone1-1.yaml
sleep 5
kubectl get pods -o wide

kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"
kubectl exec "$(kubectl get pod -l app=sleep -o jsonpath='{.items[0].metadata.name}')" -c sleep -- curl -sSL "http://helloworld:5000/hello"

echo ""
echo "################################################"
echo "测试完毕，清理"
kubectl delete -f sleep-c1-w2.yaml
bash stopapp.sh
