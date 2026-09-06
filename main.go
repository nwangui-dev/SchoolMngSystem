package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nwangui-dev/SchoolMngSystem/app" // Fix this line
)

var (
	h       = flag.Bool("h", false, "show help")
	showVer = flag.Bool("v", false, "show version")
	initcfg = flag.Bool("initcfg", false, "write default config file")
	dbinit  = flag.Bool("db", false, "initialize database and run migrations")
)

func main() {
	flag.Parse()
	handleFlags()
}

func handleFlags() {
	switch {
	case *h:
		printHelp()
	case *showVer:
		printVersion()
	case *initcfg:
		if err := app.WriteDefaultConfig("/etc/school_api.yml"); err != nil {
			log.Fatalf("Init config failed: %v", err)
		}
	case *dbinit:
		if err := app.InitDatabase(); err != nil {
			log.Fatalf("DB init failed: %v", err)
		}
	default:
		if err := app.Start(); err != nil {
			log.Fatalf("App start failed: %v", err)
		}
	}
}

func printHelp() {
	fmt.Fprintln(os.Stderr, "School Management System API Service")
	fmt.Fprintln(os.Stderr, "Usage:")
	flag.PrintDefaults()
	os.Exit(0)
}

func printVersion() {
	fmt.Println("School API v1.0.0")
	os.Exit(0)
}
