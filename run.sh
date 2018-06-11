#!/usr/bin/env bash

RANCHER_DEPLOYER=$(mktemp)

curl -L https://github.com/niranjan94/rancher-deployer/releases/download/1.0.0/rancher-deployer_$(uname -s | tr '[:upper:]' '[:lower:]')_$(uname -m) -o /usr/local/bin/rancher-deployer
chmod +x ${RANCHER_DEPLOYER}

/usr/local/bin/rancher-deployer $@