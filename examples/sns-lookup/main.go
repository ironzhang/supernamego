package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ironzhang/tlog"
	"github.com/ironzhang/tlog/iface"

	"github.com/ironzhang/supernamego"
)

type options struct {
	Forever     bool
	Interval    time.Duration
	RouteParams string
	LogLevel    string
}

func (p *options) Setup() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: sns-lookup [OPTIONS] DOMAINS\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\n")
		fmt.Fprintf(flag.CommandLine.Output(), `Example: sns-lookup -route-params="X-Lane-Cluster=sim001,X-Base-Cluster=sim000" sns/https.nginx`)
		fmt.Fprintf(flag.CommandLine.Output(), "\n")
	}

	flag.BoolVar(&p.Forever, "forever", false, "loop forever")
	flag.DurationVar(&p.Interval, "interval", time.Second, "loop interval")
	flag.StringVar(&p.RouteParams, "route-params", "", "route params")
	flag.StringVar(&p.LogLevel, "log-level", "fatal", "log level")
	flag.Parse()

	if flag.NArg() <= 0 {
		flag.Usage()
		os.Exit(0)
	}
}

func setLogLevel(s string) {
	gsl, ok := tlog.GetLogger().(iface.GetSetLevel)
	if ok {
		lv, _ := iface.StringToLevel(s)
		gsl.SetLevel(lv)
	}
}

func parseRouteParams(s string) (map[string]string, error) {
	if s == "" {
		return nil, nil
	}

	m := make(map[string]string)
	tags := strings.Split(s, ",")
	for _, tag := range tags {
		keyvalues := strings.Split(tag, "=")
		if len(keyvalues) != 2 {
			return nil, fmt.Errorf("%s is an invalid tag", tag)
		}
		m[keyvalues[0]] = keyvalues[1]
	}
	return m, nil
}

func printError(domain string, err error) {
	fmt.Printf("domain: %s\n", domain)
	fmt.Printf("error: %q\n", err)
	fmt.Printf("\n")
}

func printAddress(domain string, cluster string, addr string) {
	fmt.Printf("domain: %s\n", domain)
	fmt.Printf("cluster: %s\n", cluster)
	fmt.Printf("address: %s\n", addr)
	fmt.Printf("\n")
}

func doLookup(routes map[string]string) {
	for _, domain := range flag.Args() {
		addr, cluster, err := supernamego.Lookup(context.Background(), domain, supernamego.SetRouteContext(routes))
		if err != nil {
			printError(domain, err)
		} else {
			printAddress(domain, cluster, addr)
		}
	}
}

func main() {
	var opts options
	opts.Setup()
	setLogLevel(opts.LogLevel)

	err := supernamego.AutoSetup()
	if err != nil {
		fmt.Printf("supernamego auto setup: %v\n", err)
		return
	}

	routes, err := parseRouteParams(opts.RouteParams)
	if err != nil {
		fmt.Printf("parse route params: %v\n", err)
		return
	}

	if !opts.Forever {
		doLookup(routes)
		return
	}

	for {
		doLookup(routes)
		time.Sleep(opts.Interval)
	}
}
