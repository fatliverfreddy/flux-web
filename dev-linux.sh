#!/bin/bash

commands=("kubectl proxy --port=8080" "kubectl port-forward svc/flux 3030:3030 -n flux" "docker-compose -f docker-compose-linux.yml up --build")

for cmd in "${commands[@]}"; do {
  echo "Process \"$cmd\" started";
  $cmd & pid=$!
  PID_LIST+=" $pid";
} done

trap "kill $PID_LIST" SIGINT

echo "Parallel processes have started";

wait $PID_LIST

echo
echo "All processes have completed";
