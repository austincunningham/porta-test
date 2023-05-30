#!/bin/bash
# need a secret with the admimurl and token called provideraccount
for ((i=1;i<=600;i++));
do 
kubectl apply -f - <<EOF
---
apiVersion: capabilities.3scale.net/v1beta1
kind: Product
metadata: 
  name: product-$i
spec: 
  name: product-$i
  providerAccountRef: 
    name: provideraccount
EOF
done


