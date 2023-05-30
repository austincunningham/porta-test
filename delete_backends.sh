#!/bin/bash

for ((i=1;i<=600;i++));
do 
  kubectl delete backend backend-$i --wait=false
done