# Rancher Deployer

[![Build Status](https://travis-ci.com/niranjan94/rancher-deployer.svg?branch=master)](https://travis-ci.com/niranjan94/rancher-deployer) 
[![GitHub (pre-)release](https://img.shields.io/github/release/niranjan94/rancher-deployer/all.svg)](https://github.com/niranjan94/rancher-deployer/releases/latest)



```
USAGE:
   rancher-deployer [global options] [arguments...]

GLOBAL OPTIONS:
   --config PATH, -c PATH                        PATH to a yaml config file
   --token TOKEN, -t TOKEN                       Override TOKEN for deployment
   --tag TAG, -T TAG                             Override the TAG for the docker image to use
   --image IMAGE-URL, -i IMAGE-URL               Override the Docker IMAGE-URL
   --environments ENVIRONMENTS, -e ENVIRONMENTS  ENVIRONMENTS to deploy to (comma-separated)
   --help, -h                                    show help
   --version, -v                                 print the version
```

#### Configuration Example

```yaml
rancherUrl: https://rancher.yourdomain.com
token: xyzzy123:abcdefghijklmnop
environments:
  dev:
    -
      project: c-12ab3:p-4c5def
      image: docker.io/redis
      tag: alpine
      namespace: api
      deployment: redis-server
    -
      project: c-12ab3:p-4c5def
      image: 111111111111.dkr.ecr.ap-west-1.amazonaws.com/dev-api-server
      tag: latest
      namespace: api
      deployment: api-server
  production:
    -
      project: c-4c5def:p-12ab3
      image: docker.io/redis
      tag: alpine
      namespace: api
      deployment: redis-server
    -
      project: c-4c5def:p-12ab3
      image: 111111111111.dkr.ecr.ap-west-1.amazonaws.com/production-api-server
      tag: latest
      namespace: api
      deployment: api-server
```

##### Override via Environment variables

_**Examples:**_

- `rancherUrl` - `DEPLOYER_RANCHERURL`
- `token` - `DEPLOYER_TOKEN`

If you would like you override tokens at an environment level,

`DEPLOYER_ENVIRONMENTS_<ENVIRONMENT>_TOKEN`

_**Example:**_

`DEPLOYER_ENVIRONMENTS_DEV_TOKEN=ATokenForDev`
