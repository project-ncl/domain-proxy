# Domain Proxy

This proxy enables build tools running within a network isolated unshare environment to download artifacts from a list
of allowed target hosts.

## Components

### Domain Proxy Server

* Runs outside unshare environment
* Listens on Unix domain socket for connections from Domain Proxy Client
* Forwards HTTP requests from Domain Proxy Client to target hosts
* Forwards HTTP responses from target hosts to Domain Proxy Client
* Optionally uses an internally configured proxy for HTTP connections to target hosts

### Domain Proxy Client

* Runs inside unshare environment
* Connects to Domain Proxy Server via Unix domain socket
* Listens on HTTP port for connections from build tools
* Forwards HTTP requests from build tools to Domain Proxy Server
* Forwards HTTP responses from Domain Proxy Server to build tools

## Environment variables

### Common

* DOMAIN_PROXY_BYTE_BUFFER_SIZE
    * Default: `32768`
* DOMAIN_PROXY_DOMAIN_SOCKET
    * Default: `/tmp/domain-socket.sock`
* DOMAIN_PROXY_CONNECTION_TIMEOUT
    * Default: `10000` milliseconds
* DOMAIN_PROXY_IDLE_TIMEOUT
    * Default: `30000` milliseconds

### Domain Proxy Server

* DOMAIN_PROXY_TARGET_ALLOWLIST
    * Default:
      `localhost,repo.maven.apache.org,repository.jboss.org,packages.confluent.io,jitpack.io,repo.gradle.org,plugins.gradle.org`
* DOMAIN_PROXY_ENABLE_INTERNAL_PROXY
    * Default: `false`
* DOMAIN_PROXY_INTERNAL_PROXY_HOST
    * Default: `indy-generic-proxy`
* DOMAIN_PROXY_INTERNAL_PROXY_PORT
    * Default: `80`
* DOMAIN_PROXY_INTERNAL_PROXY_USER
* DOMAIN_PROXY_INTERNAL_PROXY_PASSWORD
* DOMAIN_PROXY_INTERNAL_NON_PROXY_HOSTS
    * Default: `localhost`

### Domain Proxy Client

* DOMAIN_PROXY_HTTP_PORT
    * Default: `8080`

## Build

* `go build -v -o bin/domainproxyserver cmd/server/main.go`
* `go build -v -o bin/domainproxyclient cmd/client/main.go`

## Run

* `./bin/domainproxyserver`
* `./bin/domainproxyclient`

## Test

* `go test -v ./test`
