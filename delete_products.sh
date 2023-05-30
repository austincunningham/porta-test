#!/bin/bash

for ((i=1;i<=600;i++));
do 
  kubectl delete product product-$i --wait=false
done