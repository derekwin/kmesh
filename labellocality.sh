kubectl label node c1-worker topology.kubernetes.io/region=region1
kubectl label node c1-worker topology.kubernetes.io/zone=zone1

kubectl label node c1-worker2 topology.kubernetes.io/region=region1
kubectl label node c1-worker2 topology.kubernetes.io/zone=zone2

kubectl label node c1-worker3 topology.kubernetes.io/region=region2
kubectl label node c1-worker3 topology.kubernetes.io/zone=zone3
