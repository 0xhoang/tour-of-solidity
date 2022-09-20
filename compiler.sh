#!/bin/bash

if [[ $1 == "callee" ]]; then
  solc --abi contracts/callee.sol -o build && abigen --abi=./build/Callee.abi --pkg=Callee --out=pkg/Callee.go
elif [[ $1 == "caller" ]]; then
  solc --abi contracts/caller.sol -o build &&
    abigen --abi=./build/Caller.abi --pkg=Caller --out=pkg/Caller.go
elif [[ $1 == "counter" ]]; then
  solc --abi contracts/Counter.sol -o build &&
    abigen --abi=./build/Counter.abi --pkg=Counter --out=pkg/Counter.go
elif [[ $1 == "simplestorage" ]]; then
  solc --abi contracts/SimpleStorage.sol -o build &&
    abigen --abi=./build/SimpleStorage.abi --pkg=SimpleStorage --out=pkg/SimpleStorage.go
else
  echo Hey that\'s is wrong params.
fi
