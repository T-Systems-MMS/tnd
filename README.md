# Trusted Network Detection

This repository contains a Trusted Network Detection (TND) implementation in
Go. It probes trusted HTTPS servers that are only reachable on a trusted
network and compares their fingerprints with predetermined values to detect if
the host is connected to a trusted network.

The TND periodically probes the trusted HTTPS servers. It detects changes to
the host's routing table and to `resolv.conf` files and triggers additional
probes in these cases. The user can retrieve the probing results from a results
channel.

## Usage

You can use the Trusted Network Detection as shown in the following example:

```golang
package main

import "github.com/T-Systems-MMS/tnd/pkg/trustnet"

func main() {
	// create tnd
	tnd := trustnet.NewTND()

	// set trusted https server(s)
	url := "https://trusted1.mynetwork.com:443"
	hash := "ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789"
	tnd.AddServer(url, hash)

	// start tnd
	tnd.Start()
	for r := range tnd.Results() {
		log.Println("Trusted Network:", r)
	}
}
```

See [examples/tnd/main.go](examples/tnd/main.go) and
[scripts/tnd.sh](scripts/tnd.sh) for a complete example.
