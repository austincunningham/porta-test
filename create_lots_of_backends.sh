#!/bin/bash
# need a secret with the admimurl and token called provideraccount


for ((i=1;i<=600;i++));
do 
kubectl apply -f - <<EOF
---
apiVersion: capabilities.3scale.net/v1beta1
kind: Backend
metadata: 
  name: backend-$i
spec: 
  name: backend-$i
  systemName: backend-$i
  privateBaseURL: "https://api.example.com"
  providerAccountRef: 
    name: provideraccount
EOF
done
