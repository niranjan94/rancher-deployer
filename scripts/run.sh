#!/usr/bin/env bash

RANCHER_DEPLOYER=$(mktemp)

curl -L https://github.com/niranjan94/rancher-deployer/releases/download/${VERSION}/rancher-deployer_$(uname -s | tr '[:upper:]' '[:lower:]')_$(uname -m) -o ${RANCHER_DEPLOYER}
chmod +x ${RANCHER_DEPLOYER}

${RANCHER_DEPLOYER} $@