#!/bin/bash
# need a secret with the admimurl and token called provideraccount

for ((i=1;i<=600;i++));
do 
kubectl apply -f - <<EOF
---
apiVersion: v1
kind: Secret
metadata:
  name: myusername-$i
stringData:
  password: "123456"
EOF

kubectl apply -f - <<EOF
---
apiVersion: capabilities.3scale.net/v1beta1
kind: DeveloperUser
metadata:
  name: developeruser-$i
  namespace: 3scale-test
spec:
  developerAccountRef:
    name: developeraccount-$i
  email: myusername-$i@example.com
  passwordCredentialsRef:
    name: myusername-$i
  providerAccountRef:
    name: provideraccount
  role: admin
  username: "myusername-$i"
EOF

kubectl apply -f - <<EOF
---
apiVersion: capabilities.3scale.net/v1beta1
kind: DeveloperAccount
metadata:
  name: developeraccount-$i
spec:
  orgName: test
  monthlyBillingEnabled : false
  monthlyChargingEnabled : false
  providerAccountRef:
    name: provideraccount
EOF
done