#!/bin/bash

for ((i=1;i<=600;i++));
do 
  #kubectl delete developerAccount developeraccount-$i --wait=false
  #kubectl delete developerUser developeruser-$i --wait=false
  kubectl delete secret myusername-$i --wait=false
done