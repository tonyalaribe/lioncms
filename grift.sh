#!/bin/bash

set -e
export GOPATH=/root/go
export PATH=/root/go/bin:/usr/local/go/bin:$PATH
cd ~/go/src/github.com/tonyalaribe/lion2018
export GO_ENV=production
/root/bin/grift $1
