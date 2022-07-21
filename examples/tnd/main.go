package main

import (
	"flag"
	"strings"

	"github.com/T-Systems-MMS/tnd/pkg/trustnet"
	log "github.com/sirupsen/logrus"
)

var (
	// parsed https servers
	httpsServers = make(map[string]string)
)

// parseCommandLine parses the command line arguments
func parseCommandLine() {
	// define and parse command line arguments
	hs := flag.String("httpsservers", "",
		"comma-separated list of trusted https server url:hash pairs")
	flag.Parse()

	// parse https servers
	if *hs == "" {
		log.Fatal("TND https servers not specified")
	}
	for _, s := range strings.Split(*hs, ",") {
		i := strings.LastIndex(s, ":")
		if i == -1 || len(s) < i+2 {
			// TODO: check a minimum hash length?
			log.Fatal("TND https server hash invalid")
		}
		url := s[:i]
		hash := strings.ToLower(s[i+1:])
		httpsServers[url] = hash
	}
}

func main() {
	// set log level
	log.SetLevel(log.DebugLevel)

	// parse command line arguments
	parseCommandLine()

	// create tnd
	tnd := trustnet.NewTND()

	// set trusted https servers
	for url, hash := range httpsServers {
		tnd.AddServer(url, hash)
	}

	// start tnd
	tnd.Start()
	for r := range tnd.Results() {
		log.WithField("trusted", r).Info("TND result")
	}
}
