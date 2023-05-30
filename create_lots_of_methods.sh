#!/bin/bash
# need a secret with the admimurl and token called provideraccount

echo '---
apiVersion: capabilities.3scale.net/v1beta1
kind: Backend
metadata: 
  name: backend
spec: 
  name: backend
  systemName: backend
  privateBaseURL: "https://api.example.com"
  providerAccountRef: 
    name: provideraccount
  methods:
    method0:
        friendlyName: method0'> backend.yaml

for ((i=1;i<=600;i++));
do 
echo "    method"$i":
        friendlyName: method"$i"" >> backend.yaml
done
