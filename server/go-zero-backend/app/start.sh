#!/bin/bash

redis-cli &
etcd &
./user/rpc/rpc -f ./user/rpc/etc/user.yaml &
./user/api/user -f ./user/api/etc/user-api.yaml &
./data/api/api -f ./data/api/etc/data-api.yaml &
./question/api/api -f ./question/api/etc/question-api.yaml &
./university/rpc/rpc -f ./university/rpc/etc/university.yaml &
./university/api/api -f ./university/api/etc/university-api.yaml
