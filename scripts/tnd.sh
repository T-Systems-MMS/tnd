#!/bin/bash

# example script with two trusted https servers

# first server
HTTP1="https://your.first.trusted.http.server.com:443"
HASH1="SHA256 hash of HTTP1"
SERV1="$HTTP1:$HASH1"

# second server
HTTP2="https://your.second.trusted.http.server.com:443"
HASH2="SHA256 hash of HTTP2"
SERV2="$HTTP2:$HASH2"

go run ./examples/tnd \
	-httpsservers "$SERV1,$SERV2"
