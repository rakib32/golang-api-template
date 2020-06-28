#! /bin/sh

# booting up dependecy containers
#docker-compose up -d

# Build go program
echo "building tests API ..."
export GO111MODULE=on
CGO_ENABLED=0 GOFLAGS=-mod=vendor go build

# setting KV, dependecy of app
echo "putting consul config ..."
curl --request PUT --data-binary @config.local.yml http://localhost:8500/v1/kv/test-api
echo "$(date "+%H:%M:%S") - consul values updated!"

# Run the app
export CONSUL_URL="127.0.0.1:8500"
export CONSUL_PATH="test-api"
./test-api serve
